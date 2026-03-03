package domain

import (
	v1 "common/api/notify/v1"
	"common/pkg/constant"
	commonModel "common/pkg/model"
	handlerchain2 "common/pkg/util/handlerchain"
	"context"
	"encoding/json"
	domainbase "notify/internal/biz/base"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"runtime/debug"

	"github.com/panjf2000/ants/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/samber/lo"
)

type EventHandler struct {
	*domainbase.BaseDomain
	handlerMap               map[string]handlerchain2.Handler[*commonModel.Notification]
	notificationTemplateRepo repo.NotificationTemplateRepo

	workerCount int
	pool        *ants.Pool
	ctx         context.Context
	cancel      context.CancelFunc
}

func NewEventHandler(base *domainbase.BaseDomain, handlerMap map[string]handlerchain2.Handler[*commonModel.Notification], notificationTemplateRepo repo.NotificationTemplateRepo) (*EventHandler, func(), error) {
	workCount := 16
	pool, err := ants.NewPool(
		workCount,
		ants.WithNonblocking(false),
		ants.WithPanicHandler(func(err interface{}) {
			// 错误兜底逻辑
			base.Log.Errorf("[ants] worker panic recovered: %v\n%s", err, debug.Stack())
		}),
	)
	if err != nil {
		return nil, nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	h := &EventHandler{
		BaseDomain:               base,
		handlerMap:               handlerMap,
		notificationTemplateRepo: notificationTemplateRepo,
		workerCount:              workCount,
		pool:                     pool,
		ctx:                      ctx,
		cancel:                   cancel,
	}
	err = h.init()
	if err != nil {
		return nil, nil, err
	}
	return h, h.CleanUp, nil
}

// init Todo 初始化默认模板
func (h *EventHandler) init() error {
	// err := ent.WithTx(h.ctx, h.db, func(tx *gen.Client) error {
	//	templateMap, err := h.notificationTemplateRepo.GetMap(h.ctx, tx, &repo.NotificationTemplateGetReq{})
	//	if err != nil {
	//		return err
	//	}
	//
	//	templates := make([]*model.NotificationTemplate, 0)
	//	h.notificationTemplateRepo.Saves(h.ctx, h.db, templates)
	// })
	// if err != nil {
	//	return err
	// }
	return nil
}

func (h *EventHandler) defaultTemplate(notificationType v1.NotificationType, channel v1.NotificationChannel) {
	switch channel {
	default:
		switch notificationType {
		case v1.NotificationType_NOTIFICATION_TYPE_USER_REGISTER:

		}
	}
}

func (h *EventHandler) Handle() {
	for range h.workerCount {
		err := h.pool.Submit(func() {
			msgs, ch, err := h.Rabbitmq.Consume(constant.QueueNotify.String())
			if err != nil {
				h.Log.Error("consume error: %v", err)
				return
			}
			defer func(ch *amqp091.Channel) {
				err := ch.Close()
				if err != nil {
					h.Log.Error("close channel failed: %v", err)
				}
			}(ch)
			for {
				select {
				case <-h.ctx.Done():
					h.Log.Info("Handle exited due to context cancel")
					return
				case msg, ok := <-msgs:
					if !ok {
						h.Log.Info("Channel closed")
						return
					}
					var err error
					h.Log.Infof("Received message: %s", string(msg.Body))

					notification := &commonModel.Notification{}
					err = json.Unmarshal(msg.Body, notification)
					if err != nil {
						h.Log.Errorf("unmarshal failed: %v", err)
						_ = msg.Nack(false, false)
						continue
					}

					var templateMap map[string]*model.NotificationTemplate
					templateMap, err = h.notificationTemplateRepo.GetCache(h.ctx, notification.Type, notification.Channels)
					if err != nil {
						h.Log.Errorf("get template failed: %v", err)
						_ = msg.Nack(false, false)
						continue
					}

					// 构建处理器链
					factory := handlerchain2.NewHandlerFactoryWithHandlers(lo.Values(h.handlerMap)...)
					// 依次按模板处理消息
					var handler handlerchain2.Handler[*commonModel.Notification]
					for _, template := range lo.Values(templateMap) {
						handler, err = factory.BuildChainByNames(template.Processors)
						if err != nil {
							break
						}
						notification.Title = template.Title
						notification.Content = template.Content
						notification.Channel = v1.NotificationChannel(template.Channel)
						_, err = handler.Handle(h.ctx, notification)
						if err != nil {
							break
						}
					}
					if err != nil {
						h.Log.Errorf("handle failed: %v", err)
						_ = msg.Nack(false, false)
						continue
					}

					// ack 消息
					if err := msg.Ack(false); err != nil {
						h.Log.Errorf("ack failed: %v", err)
					}
				}
			}
		})
		if err != nil {
			h.Log.Error("submit task failed: %v", err)
		}
	}
}

func (h *EventHandler) CleanUp() {
	h.cancel()
	h.pool.Release()
}
