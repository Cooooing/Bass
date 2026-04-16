package repo

import (
	"common/api/gen/common"
	v1 "common/api/gen/content/v1"
	commonModel "common/pkg/model"
	"content/internal/biz/model"
	"content/internal/data/ent/gen"
	"context"
)

type TagRepo interface {
	Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error)
	Saves(ctx context.Context, tx *gen.Client, tags []*model.Tag) ([]*model.Tag, error)
	Update(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error)

	GetOne(ctx context.Context, tx *gen.Client, req *TagGetReq) (*model.Tag, error)
	GetList(ctx context.Context, tx *gen.Client, req *TagGetReq) ([]*model.Tag, error)
	GetPage(ctx context.Context, tx *gen.Client, page *common.PageRequest, req *TagGetReq) ([]*model.Tag, *common.PageReply, error)
}

type TagGetReq struct {
	TagId        *int64
	TagIds       []int64
	UserId       *int64
	Name         *string
	Names        []string
	Description  *string
	Status       *v1.TagStatus
	DomainId     *int64
	ArticleCount *commonModel.Range[int32]
}
