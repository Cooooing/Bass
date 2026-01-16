package biz

type ChatSessionDomain struct {
	*BaseDomain
}

func NewChatSessionDomain(base *BaseDomain) (*ChatSessionDomain, error) {
	return &ChatSessionDomain{
		BaseDomain: base,
	}, nil
}
