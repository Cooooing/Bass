package server

import (
	"common/pkg/constant"
	commonmodel "common/pkg/model"
	"common/pkg/util"
	commonenums "common/proto/gen/common/enums"
	cerrors "common/proto/gen/common/errors"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ErrorMessageSpec struct {
	Text map[commonenums.Language]string
	Data proto.Message
}

func (m ErrorMessageSpec) Render(language commonenums.Language, data json.RawMessage) string {
	text := m.Text[language]
	if text == "" {
		text = m.Text[commonenums.Language_LANGUAGE_ZH_CN]
	}
	if text == "" || m.Data == nil {
		return text
	}
	if len(data) == 0 {
		return ""
	}
	message := proto.Clone(m.Data)
	if err := protojson.Unmarshal(data, message); err != nil {
		return ""
	}
	reflectMessage := message.ProtoReflect()
	fields := reflectMessage.Descriptor().Fields()
	args := make([]any, 0, fields.Len())
	for i := 0; i < fields.Len(); i++ {
		args = append(args, reflectMessage.Get(fields.Get(i)).Interface())
	}
	return fmt.Sprintf(text, args...)
}

type ErrorMessages map[cerrors.BusinessErrorCode]ErrorMessageSpec

func (messages ErrorMessages) Resolve(r *http.Request, code cerrors.BusinessErrorCode, data json.RawMessage) string {
	language := commonenums.Language_LANGUAGE_ZH_CN
	if r != nil {
		if user, ok := util.GetContextValue[*commonmodel.User](r.Context(), constant.CtxUserInfo); ok && user != nil && user.Language != commonenums.Language_LANGUAGE_UNSPECIFIED {
			language = user.Language
		}
	}

	message := messages[cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN]
	if value, ok := messages[code]; ok {
		message = value
	}
	if text := message.Render(language, data); text != "" {
		return text
	}
	return cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_UNKNOWN.String()
}
