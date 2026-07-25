package server

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"log/slog"

	"github.com/go-kratos/kratos/v3/config"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// HotConfigManager 管理多个配置路径的热更新订阅。
type HotConfigManager[T proto.Message] struct {
	cfg     config.Config
	root    T
	mu      sync.Mutex
	applyMu sync.Mutex
	closed  bool
	entries map[string]*hotConfigEntry
}

type hotConfigEntry struct {
	path   string
	mu     sync.RWMutex
	fields []protoHotResolvedField
}

type protoHotResolvedField struct {
	path  string
	names []protoreflect.Name
}

// BindProtoHotFields 绑定 proto 配置字段，更新时回写到传入的原始配置对象。
//
// 传入字段的地址，如 &c.Server.App、&c.Server.Jwt。
// 绑定时父 message 不能为 nil；字段在配置源中必须已存在，否则 Watch 注册会失败。
//
// 写入语义取决于字段类型：
//   - 标量字段（string/int64 等）直接赋值，与无锁读并发时理论上有数据竞争，
//     但配置改动频率极低且 x86-64 上 pointer-width 写入硬件原子，实际无影响。
//   - message 指针字段（如 *Server_Jwt）通过 proto.Clone + 指针替换更新，原子安全。
//
// 多个字段需要原子生效时，应放入同一个 message 并绑定该 message 指针，
// 这样整个 message 只触发一次回调并整体替换，避免中间态。
//
// 配置源中删除字段或变更字段类型时 Kratos Watch 不会触发回调，需业务侧兼容。
// 热更新不做语义校验（必填、范围等），由调用方自行保证配置合法性。
func (m *HotConfigManager[T]) BindProtoHotFields(
	fields ...any,
) error {
	if len(fields) == 0 {
		return nil
	}
	if m == nil {
		return errors.New("hot config manager is nil")
	}
	if m.isNilProtoMessage(m.root) {
		return errors.New("hot config root is nil")
	}

	resolved := make([]protoHotResolvedField, 0, len(fields))
	for _, field := range fields {
		item, err := m.resolveProtoHotField(field)
		if err != nil {
			return err
		}
		resolved = append(resolved, item)
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		return errors.New("hot config manager is closed")
	}

	var created []string
	for _, field := range resolved {
		if entry, ok := m.entries[field.path]; ok {
			entry.mu.Lock()
			entry.fields = append(entry.fields, field)
			entry.mu.Unlock()
			continue
		}

		entry := &hotConfigEntry{
			path:   field.path,
			fields: []protoHotResolvedField{field},
		}
		if err := m.cfg.Watch(field.path, func(_ string, _ config.Value) {
			m.applyProtoHotFields(entry)
		}); err != nil {
			for _, p := range created {
				delete(m.entries, p)
			}
			return fmt.Errorf("watch hot config %q fail: %w", field.path, err)
		}
		m.entries[field.path] = entry
		created = append(created, field.path)
	}
	return nil
}

// Close 停止所有底层配置 watcher。
func (m *HotConfigManager[T]) Close() error {
	if m == nil {
		return nil
	}

	m.mu.Lock()
	if m.closed {
		m.mu.Unlock()
		return nil
	}
	m.closed = true
	c := m.cfg
	m.mu.Unlock()

	if c == nil {
		return nil
	}
	return c.Close()
}

func (m *HotConfigManager[T]) applyProtoHotFields(
	entry *hotConfigEntry,
) {
	entry.mu.RLock()
	fields := make([]protoHotResolvedField, len(entry.fields))
	copy(fields, entry.fields)
	entry.mu.RUnlock()
	if len(fields) == 0 {
		return
	}

	m.applyMu.Lock()
	defer m.applyMu.Unlock()

	next, err := m.newProtoMessageLike(m.root)
	if err != nil {
		slog.Error("create hot config snapshot fail", "error", err)
		return
	}
	if err := m.cfg.Scan(next); err != nil {
		slog.Error("scan hot config fail", "error", err)
		return
	}

	for _, field := range fields {
		if err := m.copyProtoHotField(m.root, next, field.names); err != nil {
			slog.Error("apply hot config fail", "path", field.path, "error", err)
		}
	}
}

func (m *HotConfigManager[T]) copyProtoHotField(
	dst proto.Message,
	src proto.Message,
	names []protoreflect.Name,
) error {
	if m.isNilProtoMessage(dst) {
		return errors.New("hot config destination is nil")
	}
	if m.isNilProtoMessage(src) {
		return errors.New("hot config source is nil")
	}
	if len(names) == 0 {
		return errors.New("hot config proto path is empty")
	}

	dstMsg := dst.ProtoReflect()
	srcMsg := src.ProtoReflect()
	for i, name := range names {
		fd := dstMsg.Descriptor().Fields().ByName(name)
		if fd == nil {
			return fmt.Errorf("proto config field %q not found in %s", name, dstMsg.Descriptor().FullName())
		}
		if i == len(names)-1 {
			if fd.IsList() || fd.IsMap() {
				return fmt.Errorf("proto config field %q list or map hot reload is unsupported", fd.FullName())
			}
			if fd.Kind() == protoreflect.MessageKind || fd.Kind() == protoreflect.GroupKind {
				if !srcMsg.Has(fd) {
					dstMsg.Clear(fd)
					return nil
				}
				cloned := proto.Clone(srcMsg.Get(fd).Message().Interface())
				dstMsg.Set(fd, protoreflect.ValueOfMessage(cloned.ProtoReflect()))
				return nil
			}
			dstMsg.Set(fd, srcMsg.Get(fd))
			return nil
		}
		if fd.Kind() != protoreflect.MessageKind && fd.Kind() != protoreflect.GroupKind {
			return fmt.Errorf("proto config field %q in %s is not message", name, dstMsg.Descriptor().FullName())
		}
		if !srcMsg.Has(fd) {
			parts := make([]string, 0, i+1)
			for _, part := range names[:i+1] {
				parts = append(parts, string(part))
			}
			return fmt.Errorf("proto config parent %q is empty", strings.Join(parts, "."))
		}
		dstMsg = dstMsg.Mutable(fd).Message()
		srcMsg = srcMsg.Get(fd).Message()
	}
	return nil
}
