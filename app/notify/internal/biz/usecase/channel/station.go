package channel

import (
	"context"

	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	notifyenum "notify/internal/enum"
)

type StationUsecase struct {
	stationMessageRepo repo.NotificationStationMessageRepo
}

func NewStationUsecase(
	stationMessageRepo repo.NotificationStationMessageRepo,
) *StationUsecase {
	return &StationUsecase{
		stationMessageRepo: stationMessageRepo,
	}
}

type StationProcessReq struct {
	NotificationContext *model.NotificationContext
	Rule                *model.NotificationRule
	AccountsByUserID    map[int64]*model.UserAccount
}

func (u *StationUsecase) Process(ctx context.Context, req *StationProcessReq) (notifyenum.NotificationChannelStatus, error) {
	if req == nil || req.NotificationContext == nil || req.Rule == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	notificationContext := req.NotificationContext
	rule := req.Rule
	if rule.StationTemplate == nil {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	seen := make(map[int64]struct{}, len(notificationContext.Recipients))
	written := false
	for _, recipient := range notificationContext.Recipients {
		if recipient == nil || recipient.UserID == 0 {
			continue
		}
		if _, ok := seen[recipient.UserID]; ok {
			continue
		}
		seen[recipient.UserID] = struct{}{}
		rendered, err := rule.StationTemplate.Render(ctx, notificationContext.TemplateData)
		if err != nil {
			continue
		}
		_, err = u.stationMessageRepo.Save(ctx, &model.NotificationStationMessage{
			EventID:    notificationContext.EventID,
			EventType:  notificationContext.EventType,
			ReceiverID: recipient.UserID,
			Title:      rendered.Title,
			Content:    rendered.Content,
			Status:     notifyenum.NotificationChannelStatusSucceeded,
		})
		if err != nil {
			return notifyenum.NotificationChannelStatusInternalError, err
		}
		written = true
	}
	if !written {
		return notifyenum.NotificationChannelStatusSkipped, nil
	}
	return notifyenum.NotificationChannelStatusSucceeded, nil
}
