package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/user/v1/enum"
)

type TokenType string

const (
	TokenTypeAccess  TokenType = "access"
	TokenTypeRefresh TokenType = "refresh"
)

var TokenTypeMap = commonenum.NewMapping[TokenType, v1.TokenType](map[TokenType]commonenum.Entry[TokenType, v1.TokenType]{
	TokenTypeAccess:  {Proto: v1.TokenType_TOKEN_TYPE_ACCESS},
	TokenTypeRefresh: {Proto: v1.TokenType_TOKEN_TYPE_REFRESH},
})

func (e TokenType) String() string {
	return string(e)
}
