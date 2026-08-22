package module

import (
	"common/proto/gen/common"
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/structpb"
)

// Bootstrap 约束业务模块自己的启动配置。
type Bootstrap interface {
	proto.Message
	GetServer() *common.Server
}

// Config 保存业务模块在单体模式下使用的启动配置。
type Config[T Bootstrap] struct {
	bootstrap T
}

// RuntimeConfig 保存模块化单体共享的运行时配置。
type RuntimeConfig struct {
	Server        *common.Server
	Http          *common.HTTP
	Grpc          *common.GRPC
	Consul        *common.Consul
	Database      *common.Database
	Redis         *common.Redis
	Nats          *common.Nats
	Trace         *common.Trace
	Observability *common.Observability
	Dtm           *common.DTM
	Alert         *common.Alert
	Event         *common.Event
}

// NewConfig 使用共享基础设施和模块配置片段构造业务模块配置。
func NewConfig[T Bootstrap](runtime *RuntimeConfig, values *structpb.Struct, name string, newBootstrap func() T) (*Config[T], error) {
	bootstrap := newBootstrap()
	if err := buildModuleBootstrap(runtime, values, name, bootstrap); err != nil {
		return nil, err
	}
	return &Config[T]{bootstrap: bootstrap}, nil
}

// Bootstrap 返回模块自己的启动配置。
func (c *Config[T]) Bootstrap() T {
	return c.bootstrap
}

// Server 返回模块服务配置。
func (c *Config[T]) Server() *common.Server {
	return c.bootstrap.GetServer()
}

func (c *RuntimeConfig) serverFor(name string) *common.Server {
	if c == nil || c.Server == nil {
		return &common.Server{Name: name}
	}
	server := proto.Clone(c.Server).(*common.Server)
	server.Name = name
	return server
}

func decode(values *structpb.Struct, message proto.Message) error {
	if values == nil {
		return fmt.Errorf("module config is required")
	}
	data, err := protojson.Marshal(values)
	if err != nil {
		return fmt.Errorf("marshal module config: %w", err)
	}
	if err = (protojson.UnmarshalOptions{DiscardUnknown: false}).Unmarshal(data, message); err != nil {
		return fmt.Errorf("decode module config: %w", err)
	}
	return nil
}

func decodeField(values *structpb.Struct, field string, message proto.Message) error {
	if values == nil {
		return fmt.Errorf("module config is required")
	}
	return decode(&structpb.Struct{Fields: map[string]*structpb.Value{
		field: structpb.NewStructValue(values),
	}}, message)
}

func buildModuleBootstrap(runtime *RuntimeConfig, values *structpb.Struct, name string, bootstrap proto.Message) error {
	if err := decodeModuleValues(values, bootstrap); err != nil {
		return err
	}
	fillSharedConfig(runtime, name, bootstrap)
	return nil
}

func decodeModuleValues(values *structpb.Struct, bootstrap proto.Message) error {
	if values == nil {
		return fmt.Errorf("module config is required")
	}
	if len(values.GetFields()) == 0 {
		return nil
	}

	fields := moduleConfigFields(bootstrap)
	switch len(fields) {
	case 0:
		return fmt.Errorf("%s has no module config field", bootstrap.ProtoReflect().Descriptor().FullName())
	case 1:
		return decodeField(values, string(fields[0]), bootstrap)
	default:
		return fmt.Errorf("%s has multiple module config fields", bootstrap.ProtoReflect().Descriptor().FullName())
	}
}

func moduleConfigFields(bootstrap proto.Message) []protoreflect.Name {
	fields := bootstrap.ProtoReflect().Descriptor().Fields()
	var names []protoreflect.Name
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.Kind() != protoreflect.MessageKind {
			continue
		}
		if isSharedConfigField(field) {
			continue
		}
		names = append(names, field.Name())
	}
	return names
}

func fillSharedConfig(runtime *RuntimeConfig, name string, bootstrap proto.Message) {
	message := bootstrap.ProtoReflect()
	if runtime == nil {
		setBootstrapField(message, "server", (&RuntimeConfig{}).serverFor(name))
		return
	}
	setBootstrapField(message, "server", runtime.serverFor(name))
	setBootstrapField(message, "http", runtime.Http)
	setBootstrapField(message, "grpc", runtime.Grpc)
	setBootstrapField(message, "consul", runtime.Consul)
	setBootstrapField(message, "database", runtime.Database)
	setBootstrapField(message, "redis", runtime.Redis)
	setBootstrapField(message, "nats", runtime.Nats)
	setBootstrapField(message, "trace", runtime.Trace)
	setBootstrapField(message, "observability", runtime.Observability)
	setBootstrapField(message, "dtm", runtime.Dtm)
	setBootstrapField(message, "alert", runtime.Alert)
	setBootstrapField(message, "event", runtime.Event)
}

func isSharedConfigField(field protoreflect.FieldDescriptor) bool {
	message := field.Message()
	return message != nil && string(message.FullName()) == "common.api.app.common."+string(message.Name())
}

func setBootstrapField(bootstrap protoreflect.Message, name protoreflect.Name, value proto.Message) {
	if value == nil {
		return
	}
	field := bootstrap.Descriptor().Fields().ByName(name)
	if field == nil || field.Kind() != protoreflect.MessageKind {
		return
	}
	bootstrap.Set(field, protoreflect.ValueOfMessage(proto.Clone(value).ProtoReflect()))
}
