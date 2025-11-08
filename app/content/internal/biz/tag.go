package biz

import (
	"content/internal/biz/repo"
)

type TagDomain struct {
	*BaseDomain
	tagRepo repo.TagRepo
}

func NewTagDomain(baseDomain *BaseDomain, tagRepo repo.TagRepo) *TagDomain {
	return &TagDomain{
		BaseDomain: baseDomain,
		tagRepo:    tagRepo,
	}
}
