package usecase

import (
	"common/proto/gen/common"
	"context"
	"fmt"
	"time"

	commonClient "common/pkg/client"
	commonenum "common/pkg/enum"
	"notify/internal/biz/base"
	"notify/internal/biz/repo"
	"notify/internal/config"

	"log/slog"
)

const inboxDeadLetterScanLimit = 100

type InboxDeadLetterScanner struct {
	log         *slog.Logger
	conf        *config.Bootstrap
	inboxRepo   repo.InboxEventRepo
	alertClient *commonClient.DeadLetterAlertClient
	cancel      context.CancelFunc
}

func NewInboxDeadLetterScanner(
	logger *slog.Logger,
	conf *config.Bootstrap,
	inboxRepo repo.InboxEventRepo,
	alertClient *commonClient.DeadLetterAlertClient,
) *InboxDeadLetterScanner {
	return &InboxDeadLetterScanner{
		log:         logger,
		conf:        conf,
		inboxRepo:   inboxRepo,
		alertClient: alertClient,
	}
}

func (s *InboxDeadLetterScanner) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	go func() {
		for {
			if err := s.scan(runCtx); err != nil {
				s.log.Error(fmt.Sprintf("scan inbox dead letters failed: %v", err))
			}
			select {
			case <-runCtx.Done():
				return
			case <-time.After(s.scanInterval()):
			}
		}
	}()
	return nil
}

func (s *InboxDeadLetterScanner) Stop(_ context.Context) error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}

func (s *InboxDeadLetterScanner) scan(ctx context.Context) error {
	status := commonenum.InboxEventStatusDead
	pageResponse, err := s.inboxRepo.Page(ctx, &repo.InboxEventPageReq{Query: &repo.InboxEventQuery{Page: &base.PageRequest{Page: 1, Size: inboxDeadLetterScanLimit}, Status: &status}})
	if err != nil {
		return err
	}
	for _, row := range pageResponse.Rows {
		lastError := ""
		if row.LastError != nil {
			lastError = *row.LastError
		}
		if err := s.alertClient.Alert(ctx, s.deadLetterConf(), s.alertConf(), &commonClient.DeadLetterAlert{
			Service:   s.serviceName(),
			Source:    "inbox",
			EventID:   row.EventID,
			EventType: string(row.EventType),
			Subject:   string(row.Subject),
			Count:     row.AttemptCount,
			LastError: lastError,
			UpdatedAt: row.UpdatedAt,
		}); err != nil {
			s.log.Warn(fmt.Sprintf("alert inbox dead letter failed: event_id=%s err=%v", row.EventID, err))
		}
	}
	return nil
}

func (s *InboxDeadLetterScanner) serviceName() string {
	if s.conf != nil && s.conf.GetServer() != nil && s.conf.GetServer().GetName() != "" {
		return s.conf.GetServer().GetName()
	}
	return "notify"
}

func (s *InboxDeadLetterScanner) deadLetterConf() *common.Event_DeadLetter {
	if s.conf != nil && s.conf.GetEvent() != nil {
		return s.conf.GetEvent().GetDeadLetter()
	}
	return nil
}

func (s *InboxDeadLetterScanner) alertConf() *common.Alert {
	if s.conf != nil {
		return s.conf.GetAlert()
	}
	return nil
}

func (s *InboxDeadLetterScanner) scanInterval() time.Duration {
	if s.conf != nil && s.conf.GetEvent() != nil && s.conf.GetEvent().GetDeadLetter() != nil && s.conf.GetEvent().GetDeadLetter().GetScanInterval() != nil && s.conf.GetEvent().GetDeadLetter().GetScanInterval().AsDuration() > 0 {
		return s.conf.GetEvent().GetDeadLetter().GetScanInterval().AsDuration()
	}
	return time.Minute
}
