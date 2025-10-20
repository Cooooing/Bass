package biz

import (
	"common/pkg/util"
	"content/internal/biz/repo"

	"github.com/sony/sonyflake/v2"
)

type ArticleDomain struct {
	*BaseDomain
	domainRepo repo.Domain

	sf *sonyflake.Sonyflake
}

func NewArticleDomain(base *BaseDomain, domainRepo repo.Domain) (*ArticleDomain, error) {
	sf, err := util.NewSonyflake()
	if err != nil {
		return nil, err
	}
	return &ArticleDomain{
		BaseDomain: base,
		domainRepo: domainRepo,
		sf:         sf,
	}, nil
}
