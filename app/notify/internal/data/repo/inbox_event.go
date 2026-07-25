package repo

import (
	commonenum "common/pkg/enum"
	"common/proto/gen/common/enums"
	"context"
	"fmt"
	"notify/internal/biz/base"
	"notify/internal/biz/model"
	bizrepo "notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/inboxevent"

	utilent "common/pkg/util/ent"
)

var _ bizrepo.InboxEventRepo = (*InboxEventRepo)(nil)

type InboxEventRepo struct {
	pageNormalizer
	db *gen.Client
}

func NewInboxEventRepo(
	db *gen.Client,
) bizrepo.InboxEventRepo {
	return &InboxEventRepo{
		db: db,
	}
}

func (r *InboxEventRepo) getClient(ctx context.Context) *gen.Client {
	if c, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return c
	}
	return r.db
}

func (r *InboxEventRepo) SaveProcessing(ctx context.Context, req *bizrepo.InboxEventSaveProcessingReq) (*bizrepo.InboxEventSaveProcessingResp, error) {
	if req == nil {
		return nil, fmt.Errorf("inbox event save request is nil")
	}
	if req.EventID == "" {
		return nil, fmt.Errorf("event id is required")
	}
	eventType, ok := commonenum.EventTypeMap.ToEnum(enums.EventType(req.EventType))
	if !ok {
		return nil, fmt.Errorf("unknown event type: %s", req.EventType.String())
	}
	if _, ok := commonenum.EventSubjectMap.ToProto(req.Subject); !ok {
		return nil, fmt.Errorf("unknown event subject: %s", req.Subject)
	}

	client := r.getClient(ctx)
	save, err := client.InboxEvent.Create().
		SetEventID(req.EventID).
		SetEventType(inboxevent.EventType(eventType)).
		SetSubject(req.Subject).
		SetPayload(req.Payload).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessing)).
		SetAttemptCount(1).
		SetProcessingStartedAt(req.Now).
		Save(ctx)
	if err == nil {
		return &bizrepo.InboxEventSaveProcessingResp{
			Event:   r.inboxEvent(save),
			Claimed: true,
		}, nil
	}
	if !gen.IsConstraintError(err) {
		return nil, err
	}

	exist, getErr := client.InboxEvent.Query().
		Where(inboxevent.EventIDEQ(req.EventID)).
		Only(ctx)
	if getErr != nil {
		return nil, getErr
	}
	return &bizrepo.InboxEventSaveProcessingResp{
		Event:   r.inboxEvent(exist),
		Claimed: false,
	}, nil
}

func (r *InboxEventRepo) Get(ctx context.Context, req *bizrepo.InboxEventQuery) (*model.InboxEvent, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	row, err := query.First(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.inboxEvent(row), nil
}

func (r *InboxEventRepo) List(ctx context.Context, req *bizrepo.InboxEventQuery) ([]*model.InboxEvent, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	list, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.InboxEvent, 0, len(list))
	for _, row := range list {
		result = append(result, r.inboxEvent(row))
	}
	return result, nil
}

func (r *InboxEventRepo) Map(ctx context.Context, req *bizrepo.InboxEventQuery) (map[int64]*model.
	InboxEvent, error) {
	list, err := r.List(ctx, req)
	if err != nil {
		return nil, err
	}
	result := make(map[int64]*model.InboxEvent, len(list))
	for _, item := range list {
		result[item.ID] = item
	}
	return result, nil
}

func (r *InboxEventRepo) Count(ctx context.Context, req *bizrepo.InboxEventQuery) (int, error) {
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, req)
	count, err := query.Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *InboxEventRepo) Page(ctx context.Context, req *bizrepo.InboxEventQuery) (*bizrepo.InboxEventPageResp, error) {
	queryReq := req
	var pageReq *base.PageRequest
	if queryReq != nil {
		pageReq = queryReq.Page
	}
	page := r.normalizePage(pageReq)
	query := r.getClient(ctx).InboxEvent.Query()
	query = r.getQuery(query, queryReq)
	total, err := query.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}
	list, err := query.
		Limit(int(page.Size)).
		Offset(int((page.Page - 1) * page.Size)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.InboxEvent, 0, len(list))
	for _, row := range list {
		result = append(result, r.inboxEvent(row))
	}
	return &bizrepo.InboxEventPageResp{
		Rows: result,
		Page: &base.PageResp{
			Total: int64(total),
			Page:  page.Page,
			Size:  page.Size,
		},
	}, nil
}

func (r *InboxEventRepo) ClaimRetry(ctx context.Context, req *bizrepo.InboxEventClaimRetryReq) (bool, error) {
	if req == nil {
		return false, fmt.Errorf("inbox event claim retry request is nil")
	}
	maxRetry := req.MaxRetry
	if maxRetry <= 0 {
		maxRetry = 1
	}
	count, err := r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(req.EventID),
			inboxevent.AttemptCountLT(maxRetry),
			inboxevent.Or(
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusFailed)),
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusReceived)),
				inboxevent.And(
					inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusProcessing)),
					inboxevent.ProcessingStartedAtLTE(req.Now.Add(-req.ProcessingTimeout)),
				),
			),
		).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessing)).
		SetProcessingStartedAt(req.Now).
		AddAttemptCount(1).
		ClearLastError().
		ClearProcessedAt().
		Save(ctx)
	if err != nil || count > 0 {
		return count > 0, err
	}
	_, err = r.getClient(ctx).InboxEvent.Update().
		Where(
			inboxevent.EventIDEQ(req.EventID),
			inboxevent.AttemptCountGTE(maxRetry),
			inboxevent.Or(
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusFailed)),
				inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusReceived)),
				inboxevent.And(
					inboxevent.StatusEQ(inboxevent.Status(commonenum.InboxEventStatusProcessing)),
					inboxevent.ProcessingStartedAtLTE(req.Now.Add(-req.ProcessingTimeout)),
				),
			),
		).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusDead)).
		Save(ctx)
	return false, err
}

func (r *InboxEventRepo) MarkProcessed(ctx context.Context, req *bizrepo.InboxEventMarkProcessedReq) error {
	if req == nil {
		return fmt.Errorf("inbox event mark processed request is nil")
	}
	err := r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(req.EventID)).
		SetStatus(inboxevent.Status(commonenum.InboxEventStatusProcessed)).
		SetProcessedAt(req.Now).
		ClearLastError().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *InboxEventRepo) MarkFailed(ctx context.Context, req *bizrepo.InboxEventMarkFailedReq) error {
	if req == nil {
		return fmt.Errorf("inbox event mark failed request is nil")
	}
	maxRetry := req.MaxRetry
	if maxRetry <= 0 {
		maxRetry = 1
	}
	lastError := req.LastError
	if len(lastError) > 255 {
		lastError = lastError[:255]
	}
	event, err := r.getClient(ctx).InboxEvent.Query().
		Where(inboxevent.EventIDEQ(req.EventID)).
		Only(ctx)
	if err != nil {
		return err
	}
	status := commonenum.InboxEventStatusFailed
	if event.AttemptCount >= maxRetry {
		status = commonenum.InboxEventStatusDead
	}
	err = r.getClient(ctx).InboxEvent.Update().
		Where(inboxevent.EventIDEQ(req.EventID)).
		SetStatus(inboxevent.Status(status)).
		SetLastError(lastError).
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *InboxEventRepo) getQuery(query *gen.InboxEventQuery, req *bizrepo.InboxEventQuery) *gen.InboxEventQuery {
	if req == nil {
		return query
	}
	if req.ID != nil {
		query = query.Where(inboxevent.IDEQ(*req.ID))
	}
	if len(req.IDs) > 0 {
		query = query.Where(inboxevent.IDIn(req.IDs...))
	}
	if req.EventID != nil {
		query = query.Where(inboxevent.EventIDEQ(*req.EventID))
	}
	if len(req.EventIDs) > 0 {
		query = query.Where(inboxevent.EventIDIn(req.EventIDs...))
	}
	if req.EventType != nil {
		query = query.Where(inboxevent.EventTypeEQ(inboxevent.EventType(*req.EventType)))
	}
	if req.Subject != nil {
		query = query.Where(inboxevent.SubjectEQ(*req.Subject))
	}
	if req.Status != nil {
		query = query.Where(inboxevent.StatusEQ(inboxevent.Status(*req.Status)))
	}
	return query
}

func (r *InboxEventRepo) inboxEvent(row *gen.InboxEvent) *model.InboxEvent {
	if row == nil {
		return nil
	}
	return &model.InboxEvent{
		ID:                  row.ID,
		EventID:             row.EventID,
		EventType:           commonenum.EventType(row.EventType),
		Subject:             row.Subject,
		Payload:             row.Payload,
		Status:              commonenum.InboxEventStatus(row.Status),
		AttemptCount:        row.AttemptCount,
		LastError:           row.LastError,
		ProcessingStartedAt: row.ProcessingStartedAt,
		ProcessedAt:         row.ProcessedAt,
		CreatedAt:           row.CreatedAt,
		UpdatedAt:           row.UpdatedAt,
	}
}
