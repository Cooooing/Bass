package data

import (
	cv1 "common/api/common/v1"
	v1 "common/api/content/v1"
	"content/internal/biz/model"
	"content/internal/biz/repo"
	"content/internal/data/ent/gen"
	"content/internal/data/ent/gen/tag"
	"context"

	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TagRepo struct {
	*BaseRepo
}

func NewTagRepo(baseRepo *BaseRepo) repo.TagRepo {
	return &TagRepo{
		BaseRepo: baseRepo,
	}
}

func (t *TagRepo) Save(ctx context.Context, tx *gen.Client, tag *model.Tag) (*model.Tag, error) {
	save, err := tx.Tag.Create().
		SetUserID(tag.UserID).
		SetName(tag.Name).
		SetNillableDomainID(tag.DomainID).
		SetStatus(int32(cv1.TagStatus_TagNormal)).
		Save(ctx)
	return (*model.Tag)(save), err
}

func (t *TagRepo) GetById(ctx context.Context, tx *gen.Client, id int64) (*model.Tag, error) {
	query, err := tx.Tag.Query().Where(tag.IDEQ(id)).First(ctx)
	return (*model.Tag)(query), err
}

func (t *TagRepo) GetList(ctx context.Context, tx *gen.Client, req *v1.GetTagRequest) (*v1.GetTagReply, error) {
	query := tx.Tag.Query()
	countQuery := query.Clone()
	count, err := countQuery.Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.Limit(int(req.Page.Size)).Offset(int((req.Page.Page - 1) * req.Page.Size)).All(ctx)
	if err != nil {
		return nil, err
	}

	tags := make([]*v1.Tag, 0, len(list))
	for _, item := range list {
		i := &v1.Tag{}
		err = copier.Copy(i, item)
		if err != nil {
			return nil, err
		}
		i.CreatedAt = timestamppb.New(*item.CreatedAt)
		i.UpdatedAt = timestamppb.New(*item.UpdatedAt)
		tags = append(tags, i)
	}
	return &v1.GetTagReply{
		Page: &cv1.PageReply{
			Total: uint32(count),
			Size:  req.Page.Size,
			Page:  req.Page.Page,
		},
		Tags: tags,
	}, nil
}
