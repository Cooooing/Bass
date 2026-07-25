package apperror

import (
	cerrors "common/proto/gen/common/errors"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"

	kratoserrors "github.com/go-kratos/kratos/v3/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	metadataBusinessCode = "business_code"
	metadataData         = "data"
)

type Option func(metadata map[string]string)

func New(code cerrors.BusinessErrorCode, opts ...Option) *kratoserrors.Error {
	return NewMessage(code, "", opts...)
}

func NewMessage(code cerrors.BusinessErrorCode, message string, opts ...Option) *kratoserrors.Error {
	if !validBusinessCode(code) || code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_SUCCESS {
		code = cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN
	}
	if message == "" {
		message = code.String()
	}
	metadata := map[string]string{
		metadataBusinessCode: strconv.Itoa(int(code)),
	}
	for _, opt := range opts {
		if opt != nil {
			opt(metadata)
		}
	}
	return kratoserrors.New(StatusCode(code), code.String(), message).WithMetadata(metadata)
}

func WithData[T proto.Message](data T) Option {
	return func(metadata map[string]string) {
		value := reflect.ValueOf(data)
		if !value.IsValid() {
			return
		}
		switch value.Kind() {
		case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Slice:
			if value.IsNil() {
				return
			}
		}
		payload, err := protojson.MarshalOptions{
			UseProtoNames: true,
		}.Marshal(data)
		if err != nil {
			return
		}
		metadata[metadataData] = string(payload)
	}
}

func CommonInvalidArgument() *kratoserrors.Error {
	return New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
}

func GameTownWorldInvalid() *kratoserrors.Error {
	return New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID)
}

func GameTownWorldInvalidMessage(message string) *kratoserrors.Error {
	return NewMessage(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_GAME_TOWN_WORLD_INVALID, message)
}

func BusinessCode(err error) (cerrors.BusinessErrorCode, bool) {
	se := kratoserrors.FromError(err)
	if se == nil {
		return 0, false
	}
	if se.Metadata != nil {
		raw := se.Metadata[metadataBusinessCode]
		if raw != "" {
			value, err := strconv.ParseInt(raw, 10, 32)
			if err == nil {
				code := cerrors.BusinessErrorCode(value)
				if validBusinessCode(code) {
					return code, true
				}
			}
		}
	}
	if value, ok := cerrors.BusinessErrorCode_value[se.Reason]; ok {
		code := cerrors.BusinessErrorCode(value)
		if validBusinessCode(code) {
			return code, true
		}
	}
	return 0, false
}

func Data(err error) json.RawMessage {
	se := kratoserrors.FromError(err)
	if se == nil || se.Metadata == nil {
		return nil
	}
	raw := se.Metadata[metadataData]
	if raw == "" {
		return nil
	}
	return json.RawMessage(raw)
}

func StatusCode(code cerrors.BusinessErrorCode) int {
	if !validBusinessCode(code) {
		return http.StatusInternalServerError
	}
	value := code.Descriptor().Values().ByNumber(code.Number())
	if value != nil {
		if statusCode := extensionInt32(value.Options(), kratoserrors.E_Code); statusCode > 0 {
			return statusCode
		}
	}
	if statusCode := extensionInt32(code.Descriptor().Options(), kratoserrors.E_DefaultCode); statusCode > 0 {
		return statusCode
	}
	return http.StatusInternalServerError
}

func IsInternal(err error) bool {
	if code, ok := BusinessCode(err); ok {
		return code == cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INTERNAL
	}
	se := kratoserrors.FromError(err)
	return se != nil && se.Code >= http.StatusInternalServerError
}

func validBusinessCode(code cerrors.BusinessErrorCode) bool {
	_, ok := cerrors.BusinessErrorCode_name[int32(code)]
	return ok
}

func extensionInt32(message proto.Message, extension protoreflect.ExtensionType) int {
	if message == nil || !proto.HasExtension(message, extension) {
		return 0
	}
	switch value := proto.GetExtension(message, extension).(type) {
	case int32:
		return int(value)
	case *int32:
		if value != nil {
			return int(*value)
		}
	}
	return 0
}
