package usecase

import (
	"context"
	"time"

	commonClient "common/pkg/client"
	commonenum "common/pkg/enum"
	"common/proto/gen/common"
	"user/internal/biz/repo"
	"user/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
)

const outboxDeadLetterScanLimit = 100

type OutboxDeadLetterScanner struct {
	log         *log.Helper
	conf        *conf.Bootstrap
	outboxRepo  repo.OutboxEventRepo
	alertClient *commonClient.DeadLetterAlertClient
	cancel      context.CancelFunc
}

func NewOutboxDeadLetterScanner(
	logger log.Logger,
	conf *conf.Bootstrap,
	outboxRepo repo.OutboxEventRepo,
	alertClient *commonClient.DeadLetterAlertClient,
) *OutboxDeadLetterScanner {
	return &OutboxDeadLetterScanner{
		log:         log.NewHelper(logger),
		conf:        conf,
		outboxRepo:  outboxRepo,
		alertClient: alertClient,
	}
}

func (s *OutboxDeadLetterScanner) Start(ctx context.Context) error {
	runCtx, cancel := context.WithCancel(ctx)
	s.cancel = cancel
	go func() {
		for {
			if err := s.scan(runCtx); err != nil {
				s.log.Errorf("scan outbox dead letters failed: %v", err)
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

func (s *OutboxDeadLetterScanner) Stop(_ context.Context) error {
	if s.cancel != nil {
		s.cancel()
	}
	return nil
}

func (s *OutboxDeadLetterScanner) scan(ctx context.Context) error {
	status := commonenum.OutboxEventStatusDead
	rows, _, err := s.outboxRepo.Page(ctx, &common.PageRequest{Page: 1, Size: outboxDeadLetterScanLimit}, &repo.OutboxEventGetReq{Status: &status})
	if err != nil {
		return err
	}
	for _, row := range rows {
		lastError := ""
		if row.LastError != nil {
			lastError = *row.LastError
		}
		if err := s.alertClient.Alert(ctx, s.deadLetterConf(), s.alertConf(), &commonClient.DeadLetterAlert{
			Service:   s.serviceName(),
			Source:    "outbox",
			EventID:   row.EventID,
			EventType: string(row.EventType),
			Subject:   string(row.Subject),
			Count:     row.RetryCount,
			LastError: lastError,
			UpdatedAt: row.UpdatedAt,
		}); err != nil {
			s.log.Warnf("alert outbox dead letter failed: event_id=%s err=%v", row.EventID, err)
		}
	}
	return nil
}

func (s *OutboxDeadLetterScanner) serviceName() string {
	if s.conf != nil && s.conf.GetServer() != nil && s.conf.GetServer().GetName() != "" {
		return s.conf.GetServer().GetName()
	}
	return "user"
}

func (s *OutboxDeadLetterScanner) deadLetterConf() *common.Event_DeadLetter {
	if s.conf != nil && s.conf.GetEvent() != nil {
		return s.conf.GetEvent().GetDeadLetter()
	}
	return nil
}

func (s *OutboxDeadLetterScanner) alertConf() *common.Alert {
	if s.conf != nil {
		return s.conf.GetAlert()
	}
	return nil
}

func (s *OutboxDeadLetterScanner) scanInterval() time.Duration {
	if s.conf != nil && s.conf.GetEvent() != nil && s.conf.GetEvent().GetDeadLetter() != nil && s.conf.GetEvent().GetDeadLetter().GetScanInterval() != nil && s.conf.GetEvent().GetDeadLetter().GetScanInterval().AsDuration() > 0 {
		return s.conf.GetEvent().GetDeadLetter().GetScanInterval().AsDuration()
	}
	return time.Minute
}
