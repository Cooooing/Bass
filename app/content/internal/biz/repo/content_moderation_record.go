package repo

import (
	"context"

	"content/internal/biz/base"
	"content/internal/biz/model"
	"content/internal/enum"
)

type ContentModerationRecordRepo interface {
	Save(ctx context.Context, req *ContentModerationRecordSaveReq) (*ContentModerationRecordSaveResponse, error)
	Get(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordGetResponse, error)
	List(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordListResponse, error)
	Map(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordMapResponse, error)
	Count(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordCountResponse, error)
	Page(ctx context.Context, req *ContentModerationRecordGetReq) (*ContentModerationRecordPageResponse, error)
}

type ContentModerationRecordSaveReq struct {
	Record *model.ContentModerationRecord
}

type ContentModerationRecordSaveResponse struct {
	Record *model.ContentModerationRecord
}

type ContentModerationRecordGetResponse struct {
	Record *model.ContentModerationRecord
}

type ContentModerationRecordListResponse struct {
	Rows []*model.ContentModerationRecord
}

type ContentModerationRecordMapResponse struct {
	Rows map[int64]*model.ContentModerationRecord
}

type ContentModerationRecordCountResponse struct {
	Count int
}

type ContentModerationRecordPageResponse struct {
	Rows []*model.ContentModerationRecord
	Page *base.PageResponse
}

type ContentModerationRecordGetReq struct {
	Page       *base.PageRequest
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
