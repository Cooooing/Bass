package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	realtimehubv1 "common/proto/gen/realtime_hub/v1"
	"realtime_node/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	kratoshttp "github.com/go-kratos/kratos/v3/transport/http"
	"github.com/google/uuid"
)

type RealtimeService struct {
	realtime *usecase.RealtimeUsecase
}

func NewRealtimeService(realtime *usecase.RealtimeUsecase) *RealtimeService {
	return &RealtimeService{realtime: realtime}
}

func (*RealtimeService) RegisterGrpc(*grpc.Server) {}

func (s *RealtimeService) RegisterHTTP(server *kratoshttp.Server) {
	server.Handle("/v1/stream/bbs/sse", http.HandlerFunc(s.HandleSSE))
}

func (s *RealtimeService) HandleSSE(writer http.ResponseWriter, request *http.Request) {
	ticket := request.URL.Query().Get("ticket")
	if ticket == "" {
		http.Error(writer, "unauthorized", http.StatusUnauthorized)
		return
	}
	userID, err := s.realtime.Authenticate(request.Context(), ticket)
	if err != nil {
		http.Error(writer, "unauthorized", http.StatusUnauthorized)
		return
	}
	flusher, ok := writer.(http.Flusher)
	if !ok {
		http.Error(writer, "stream unavailable", http.StatusInternalServerError)
		return
	}
	connectionID := uuid.NewString()
	deliveries := make(chan *realtimehubv1.Delivery, 32)
	connectionCtx, cancel := context.WithCancel(request.Context())
	defer cancel()
	if err := s.realtime.Connect(connectionCtx, userID, connectionID, deliveries, cancel); err != nil {
		http.Error(writer, "unavailable", http.StatusServiceUnavailable)
		return
	}
	defer s.realtime.Disconnect(userID, connectionID)
	writer.Header().Set("Content-Type", "text/event-stream")
	writer.Header().Set("Cache-Control", "no-cache")
	writer.Header().Set("Referrer-Policy", "no-referrer")
	if _, err := fmt.Fprint(writer, ":connected\n\n"); err != nil {
		return
	}
	flusher.Flush()
	ticker := time.NewTicker(25 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-connectionCtx.Done():
			return
		case delivery := <-deliveries:
			event, ok := s.realtime.ClientEvent(delivery)
			if !ok {
				continue
			}
			data, err := json.Marshal(event.Payload)
			if err != nil {
				continue
			}
			if _, err := fmt.Fprintf(writer, "event: %s\ndata: %s\n\n", event.Name, data); err != nil {
				return
			}
			flusher.Flush()
			if err := s.realtime.RenewConnection(connectionCtx, userID, connectionID); err != nil {
				return
			}
		case <-ticker.C:
			if _, err := fmt.Fprint(writer, ":keepalive\n\n"); err != nil {
				return
			}
			flusher.Flush()
			if err := s.realtime.RenewConnection(connectionCtx, userID, connectionID); err != nil {
				return
			}
		}
	}
}
