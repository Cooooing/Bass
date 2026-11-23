package server

import (
	event "common/proto/gen/common/enums/event/v1"
	"context"

	"notify/internal/biz/usecase"
)

type TemplateInitializationServer struct {
	templateUsecase *usecase.TemplateUsecase
	eventHandlers   map[event.EventSubject][]usecase.EventHandler
}

func NewTemplateInitializationServer(
	templateUsecase *usecase.TemplateUsecase,
	eventHandlers map[event.EventSubject][]usecase.EventHandler,
) *TemplateInitializationServer {
	return &TemplateInitializationServer{
		templateUsecase: templateUsecase,
		eventHandlers:   eventHandlers,
	}
}

func (s *TemplateInitializationServer) Start(ctx context.Context) error {
	if s == nil || s.templateUsecase == nil {
		return nil
	}
	return s.templateUsecase.InitDefaultTemplates(ctx, s.eventHandlers)
}

func (s *TemplateInitializationServer) Stop(ctx context.Context) error {
	return nil
}
