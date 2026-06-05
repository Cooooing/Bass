package server

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
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
func (m *HotConfigManager[T]) BindProtoHotFields(fields ...any) error {
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
			return fmt.Errorf("watch hot config %q fail: %w", field.path, err)
		}
		m.entries[field.path] = entry
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

func (m *HotConfigManager[T]) applyProtoHotFields(entry *hotConfigEntry) {
	entry.mu.RLock()
	fields := make([]protoHotResolvedField, len(entry.fields))
	copy(fields, entry.fields)
	entry.mu.RUnlock()
	if len(fields) == 0 {
		return
	}

	m.applyMu.Lock()
	defer m.applyMu.Unlock()
	for _, field := range fields {
		next, err := m.newProtoMessageLike(m.root)
		if err != nil {
			log.Errorf("create hot config %s snapshot fail: %v", field.path, err)
			continue
		}
		if err := m.cfg.Scan(next); err != nil {
			log.Errorf("scan hot config %s fail: %v", field.path, err)
			continue
		}
		if err := m.copyProtoHotField(m.root, next, field.names); err != nil {
			log.Errorf("apply hot config %s fail: %v", field.path, err)
		}
	}
}

func (m *HotConfigManager[T]) copyProtoHotField(dst proto.Message, src proto.Message, names []protoreflect.Name) error {
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
