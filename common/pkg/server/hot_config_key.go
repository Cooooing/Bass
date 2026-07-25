package server

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var protoMessageType = reflect.TypeOf((*proto.Message)(nil)).Elem()

func (m *HotConfigManager[T]) resolveProtoHotField(
	field any,
) (protoHotResolvedField, error) {
	if m.isNilProtoMessage(m.root) {
		return protoHotResolvedField{}, errors.New("hot config root is nil")
	}

	target := reflect.ValueOf(field)
	if !target.IsValid() {
		return protoHotResolvedField{}, errors.New("hot config field is nil")
	}
	if target.Kind() != reflect.Pointer {
		return protoHotResolvedField{}, fmt.Errorf("hot config field must be field address, got %s", target.Kind())
	}
	if target.IsNil() {
		return protoHotResolvedField{}, errors.New("hot config field is nil")
	}

	parts, ok := m.findProtoHotPath(reflect.ValueOf(m.root).Elem(), target.Pointer(), nil)
	if !ok {
		return protoHotResolvedField{}, errors.New("hot config field must be address of generated proto field")
	}
	names := make([]protoreflect.Name, 0, len(parts))
	for _, part := range parts {
		names = append(names, protoreflect.Name(part))
	}
	if err := m.validateProtoHotPath(m.root, names); err != nil {
		return protoHotResolvedField{}, err
	}

	return protoHotResolvedField{
		path:  strings.Join(parts, "."),
		names: names,
	}, nil
}

func (m *HotConfigManager[T]) newProtoMessageLike(
	root proto.Message,
) (proto.Message, error) {
	if m.isNilProtoMessage(root) {
		return nil, errors.New("hot config proto root is nil")
	}
	typ := reflect.TypeOf(root)
	if typ.Kind() != reflect.Pointer || typ.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("hot config proto root must be pointer to struct: %s", typ)
	}
	msg, ok := reflect.New(typ.Elem()).Interface().(proto.Message)
	if !ok {
		return nil, fmt.Errorf("hot config proto root %s does not implement proto.Message", typ)
	}
	return msg, nil
}

func (m *HotConfigManager[T]) isNilProtoMessage(
	msg proto.Message,
) bool {
	v := reflect.ValueOf(msg)
	if !v.IsValid() {
		return true
	}
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
		return v.IsNil()
	default:
		return false
	}
}

func (m *HotConfigManager[T]) findProtoHotPath(
	v reflect.Value,
	target uintptr,
	prefix []string,
) ([]string, bool) {
	if v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, false
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, false
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		sf := typ.Field(i)
		name, ok := m.protoFieldName(sf)
		if !ok {
			continue
		}
		fv := v.Field(i)
		path := append(append([]string{}, prefix...), name)
		if fv.CanAddr() && fv.Addr().Pointer() == target {
			return path, true
		}
		if fv.Kind() == reflect.Pointer && fv.Type().Implements(protoMessageType) && !fv.IsNil() {
			if found, ok := m.findProtoHotPath(fv, target, path); ok {
				return found, true
			}
		}
	}
	return nil, false
}

func (m *HotConfigManager[T]) protoFieldName(
	field reflect.StructField,
) (string, bool) {
	tag := field.Tag.Get("protobuf")
	if tag == "" {
		return "", false
	}
	for _, part := range strings.Split(tag, ",") {
		if name, ok := strings.CutPrefix(part, "name="); ok && name != "" {
			return name, true
		}
	}
	return "", false
}

func (m *HotConfigManager[T]) validateProtoHotPath(
	root proto.Message,
	names []protoreflect.Name,
) error {
	if len(names) == 0 {
		return errors.New("hot config proto path is empty")
	}
	md := root.ProtoReflect().Descriptor()
	for i, name := range names {
		fd := md.Fields().ByName(name)
		if fd == nil {
			return fmt.Errorf("proto config field %q not found in %s", name, md.FullName())
		}
		if fd.IsList() || fd.IsMap() {
			return fmt.Errorf("proto config field %q list or map hot reload is unsupported", fd.FullName())
		}
		if i == len(names)-1 {
			return nil
		}
		if fd.Kind() != protoreflect.MessageKind && fd.Kind() != protoreflect.GroupKind {
			return fmt.Errorf("proto config field %q in %s is not message", name, md.FullName())
		}
		md = fd.Message()
	}
	return nil
}
