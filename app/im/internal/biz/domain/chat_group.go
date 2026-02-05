package domain

import domainbase "im/internal/biz/base"

type ChatGroupDomain struct {
	*domainbase.BaseDomain
}

func NewChatGroupDomain(base *domainbase.BaseDomain) (*ChatGroupDomain, error) {
	return &ChatGroupDomain{
		BaseDomain: base,
	}, nil
}
