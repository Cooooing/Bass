package repo

import (
	commonmodel "common/pkg/model"
	"content/internal/biz/model"
	"content/internal/enum"
	"context"
)

type ContentModerationRecordRepo interface {
	Save(ctx context.Context, record *model.ContentModerationRecord) (*model.ContentModerationRecord, error)
	Get(ctx context.Context, req *ContentModerationRecordGetReq) (*model.ContentModerationRecord, error)
	List(ctx context.Context, req *ContentModerationRecordGetReq) ([]*model.ContentModerationRecord, error)
	Map(ctx context.Context, req *ContentModerationRecordGetReq) (map[int64]*model.ContentModerationRecord, error)
	Count(ctx context.Context, req *ContentModerationRecordGetReq) (int, error)
	Page(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordPageResp, error)
}

type ContentModerationRecordPageResp struct {
	Rows []*model.ContentModerationRecord
	Page *commonmodel.PageResp
}

type ContentModerationRecordGetReq struct {
	Page       *commonmodel.PageReq
	ID         *int64
	IDs        []int64
	Target     *enum.ContentModerationTarget
	Targets    []enum.ContentModerationTarget
	TargetID   *int64
	TargetIDs  []int64
	Action     *enum.ContentModerationAction
	Actions    []enum.ContentModerationAction
	OperatorID *int64
}
