package server

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	kratoserrors "github.com/go-kratos/kratos/v3/errors"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtoJSONRequestDecoder(r *http.Request, v any) error {
	message, ok := v.(proto.Message)
	if !ok {
		return kratoshttp.DefaultRequestDecoder(r, v)
	}
	codec, ok := kratoshttp.CodecForRequest(r, "Content-Type")
	if !ok || codec.Name() != "json" {
		return kratoshttp.DefaultRequestDecoder(r, v)
	}
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	if err != nil {
		return kratoserrors.BadRequest("CODEC", err.Error())
	}
	if len(data) == 0 {
		return nil
	}
	if err := (protojson.UnmarshalOptions{DiscardUnknown: true}).Unmarshal(data, message); err != nil {
		return kratoserrors.BadRequest("CODEC", fmt.Sprintf("body unmarshal %s", err.Error()))
	}
	return nil
}
