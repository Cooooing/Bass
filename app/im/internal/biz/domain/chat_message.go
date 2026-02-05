package domain

import domainbase "im/internal/biz/base"

type ChatMessageDomain struct {
	*domainbase.BaseDomain
}

func NewChatMessageDomain(base *domainbase.BaseDomain) (*ChatMessageDomain, error) {
	return &ChatMessageDomain{
		BaseDomain: base,
	}, nil
}
