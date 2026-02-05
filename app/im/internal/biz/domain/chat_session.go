package domain

import domainbase "im/internal/biz/base"

type ChatSessionDomain struct {
	*domainbase.BaseDomain
}

func NewChatSessionDomain(base *domainbase.BaseDomain) (*ChatSessionDomain, error) {
	return &ChatSessionDomain{
		BaseDomain: base,
	}, nil
}
