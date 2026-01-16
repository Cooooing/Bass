package biz

type ChatMessageDomain struct {
	*BaseDomain
}

func NewChatMessageDomain(base *BaseDomain) (*ChatMessageDomain, error) {
	return &ChatMessageDomain{
		BaseDomain: base,
	}, nil
}
