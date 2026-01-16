package biz

type ChatGroupDomain struct {
	*BaseDomain
}

func NewChatGroupDomain(base *BaseDomain) (*ChatGroupDomain, error) {
	return &ChatGroupDomain{
		BaseDomain: base,
	}, nil
}
