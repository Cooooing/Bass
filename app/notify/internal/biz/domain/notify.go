package domain

import (
	"bytes"
	"common/api/gen/common/enums"
	v1 "common/api/gen/notify/v1"
	"common/pkg/client/rpc"
	"context"
	"html/template"
	"notify/internal/biz/domain/sender"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/data/gen"
	"notify/internal/data/gen/notificationmeta"
	notifyenum "notify/internal/enum"

	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/protobuf/proto"
)

// UserResolver 用户信息解析接口（由外部实现，RPC 获取手机号/邮箱等）
type UserResolver interface {
	Resolve(ctx context.Context, userID int64) (*sender.UserInfo, error)
}

// DeliveryRequest 投递请求
type DeliveryRequest struct {
	EventId     string
	EventType   enums.EventType
	ReceiverIDs []int64
	Vars        any
}

// NotifyService 纯投递服务
type NotifyService struct {
	log                      *log.Helper
	db                       *gen.Client
	notificationMetaRepo     repo.NotificationMetaRepo
	notificationRecordRepo   repo.NotificationRecordRepo
	notificationTemplateRepo repo.NotificationTemplateRepo
	notificationSettingRepo  repo.NotificationSettingRepo
	userClient               *rpc.UserClient
	userResolver             UserResolver
	senders                  *sender.Registry
}

func NewNotifyService(
	logger log.Logger,
	db *gen.Client,
	notificationMetaRepo repo.NotificationMetaRepo,
	notificationRecordRepo repo.NotificationRecordRepo,
	notificationTemplateRepo repo.NotificationTemplateRepo,
	notificationSettingRepo repo.NotificationSettingRepo,
	userClient *rpc.UserClient,
	userResolver UserResolver,
	senders *sender.Registry,
) *NotifyService {
	return &NotifyService{
		log:                      log.NewHelper(logger),
		db:                       db,
		notificationMetaRepo:     notificationMetaRepo,
		notificationRecordRepo:   notificationRecordRepo,
		notificationTemplateRepo: notificationTemplateRepo,
		notificationSettingRepo:  notificationSettingRepo,
		userClient:               userClient,
		userResolver:             userResolver,
		senders:                  senders,
	}
}

// Deliver 投递通知
func (s *NotifyService) Deliver(ctx context.Context, req *DeliveryRequest) error {
	if len(req.ReceiverIDs) == 0 {
		return nil
	}

	// 1. 序列化 payload
	var payloadBytes []byte
	if pm, ok := req.Vars.(proto.Message); ok {
		payloadBytes, _ = proto.Marshal(pm)
	}

	// 2. 创建 notification_meta（uuid 唯一约束保证幂等）
	dbEventType, _ := notifyenum.EventTypeMap.ToEnum(req.EventType)
	meta, err := s.notificationMetaRepo.Save(ctx, s.db, &model.NotificationMeta{
		NotificationMeta: &gen.NotificationMeta{
			UUID:      req.EventId,
			EventType: notificationmeta.EventType(dbEventType),
			Meta:      payloadBytes,
			Status:    notificationmeta.StatusNormal,
		},
	})
	if err != nil {
		s.log.Infof("duplicate event or save failed, skip: event_id=%s err=%v", req.EventId, err)
		return nil
	}

	// 3. 逐接收者投递（每个用户可能有不同语言，模板在投递时按语言查询）
	for _, receiverID := range req.ReceiverIDs {
		s.deliverToReceiver(ctx, req, receiverID, meta.ID)
	}

	return nil
}

func (s *NotifyService) deliverToReceiver(
	ctx context.Context,
	req *DeliveryRequest,
	receiverID int64,
	metaID int64,
) {
	// 1. 查用户语言
	language := s.getLanguage(ctx, receiverID)

	// 2. 查该语言的模板
	templates, err := s.notificationTemplateRepo.GetTemplates(ctx, req.EventType, language)
	if err != nil {
		s.log.Errorf("get templates failed: type=%v lang=%s err=%v", req.EventType, language, err)
		return
	}
	if len(templates) == 0 {
		s.log.Warnf("no templates found: type=%v lang=%s", req.EventType, language)
		return
	}

	// 3. 渲染模板
	rendered := make(map[v1.NotificationChannel]struct{ Title, Content string })
	for _, tpl := range templates {
		ch := sender.ChannelToProto(string(tpl.Channel))
		rendered[ch] = struct{ Title, Content string }{
			Title:   s.render(tpl.Title, req.Vars),
			Content: s.render(tpl.Content, req.Vars),
		}
	}

	// 4. 查用户偏好
	prefs, err := s.getUserPrefs(ctx, receiverID, req.EventType)
	if err != nil {
		s.log.Errorf("get user prefs failed: user=%d err=%v", receiverID, err)
	}

	// 5. 确定可用渠道
	enabledChannels := make(map[v1.NotificationChannel]bool)
	for ch := range rendered {
		enabledChannels[ch] = true
	}
	if len(prefs) > 0 {
		for ch := range enabledChannels {
			enabledChannels[ch] = false
		}
		for _, p := range prefs {
			if p.Enable {
				enabledChannels[sender.ChannelToProto(string(p.Channel))] = true
			}
		}
	}

	// 6. 获取用户联系信息
	var userInfo sender.UserInfo
	if s.userResolver != nil {
		info, err := s.userResolver.Resolve(ctx, receiverID)
		if err != nil {
			s.log.Errorf("resolve user info failed: user=%d err=%v", receiverID, err)
		} else if info != nil {
			userInfo = *info
		}
	}

	// 7. 按渠道发送
	for ch, enabled := range enabledChannels {
		if !enabled {
			continue
		}
		r := rendered[ch]
		if err := s.senders.Send(ctx, ch, &sender.SendRequest{
			ReceiverID: receiverID,
			Title:      r.Title,
			Content:    r.Content,
			UserInfo:   userInfo,
		}); err != nil {
			s.log.Errorf("send failed: channel=%v receiver=%d err=%v", ch, receiverID, err)
		}
	}

	// 8. 创建 notification_record
	_, err = s.notificationRecordRepo.Save(ctx, s.db, &model.NotificationRecord{
		NotificationRecord: &gen.NotificationRecord{
			NotificationID: metaID,
			ReceiverID:     receiverID,
		},
	})
	if err != nil {
		s.log.Errorf("create notification record failed: receiver=%d err=%v", receiverID, err)
	}
}

func (s *NotifyService) getUserPrefs(ctx context.Context, userID int64, eventType enums.EventType) ([]*model.NotificationSetting, error) {
	return s.notificationSettingRepo.GetByUserAndEvent(ctx, userID, eventType)
}

func (s *NotifyService) render(tplStr string, variables any) string {
	tpl, err := template.New("").Parse(tplStr)
	if err != nil {
		s.log.Errorf("parse template failed: %v", err)
		return tplStr
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, variables); err != nil {
		s.log.Errorf("execute template failed: %v", err)
		return tplStr
	}
	return buf.String()
}
