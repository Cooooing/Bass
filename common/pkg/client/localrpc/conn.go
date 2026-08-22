package localrpc

import (
	"context"
	"fmt"
	"io"
	"reflect"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

type Conn struct {
	methods map[string]method
	streams map[string]stream
}

type method struct {
	server  any
	handler grpc.MethodHandler
}

type stream struct {
	server any
	desc   grpc.StreamDesc
}

func NewConn() *Conn {
	return &Conn{
		methods: make(map[string]method),
		streams: make(map[string]stream),
	}
}

func (c *Conn) RegisterMatching(desc *grpc.ServiceDesc, server any) bool {
	serverType := reflect.TypeOf(server)
	if serverType == nil {
		return false
	}
	handlerType := reflect.TypeOf(desc.HandlerType)
	if handlerType == nil || handlerType.Kind() != reflect.Pointer {
		return false
	}
	if !serverType.Implements(handlerType.Elem()) {
		return false
	}
	c.RegisterService(desc, server)
	return true
}

func (c *Conn) RegisterService(desc *grpc.ServiceDesc, server any) {
	for _, m := range desc.Methods {
		fullMethod := fmt.Sprintf("/%s/%s", desc.ServiceName, m.MethodName)
		c.methods[fullMethod] = method{server: server, handler: m.Handler}
	}
	for _, s := range desc.Streams {
		fullMethod := fmt.Sprintf("/%s/%s", desc.ServiceName, s.StreamName)
		c.streams[fullMethod] = stream{server: server, desc: s}
	}
}

func (c *Conn) Invoke(ctx context.Context, name string, args any, reply any, _ ...grpc.CallOption) error {
	localMethod, ok := c.methods[name]
	if !ok {
		return status.Errorf(codes.Unimplemented, "local method %s not mounted", name)
	}
	req, ok := args.(proto.Message)
	if !ok {
		return status.Errorf(codes.InvalidArgument, "local method %s request is not proto message", name)
	}
	resp, ok := reply.(proto.Message)
	if !ok {
		return status.Errorf(codes.Internal, "local method %s response is not proto message", name)
	}

	out, err := localMethod.handler(localMethod.server, ctx, func(v any) error {
		dst, ok := v.(proto.Message)
		if !ok {
			return status.Errorf(codes.Internal, "local method %s decoder target is not proto message", name)
		}
		proto.Reset(dst)
		proto.Merge(dst, req)
		return nil
	}, nil)
	if err != nil {
		return err
	}
	msg, ok := out.(proto.Message)
	if !ok {
		return status.Errorf(codes.Internal, "local method %s returned non proto response", name)
	}
	proto.Reset(resp)
	proto.Merge(resp, msg)
	return nil
}

func (c *Conn) NewStream(ctx context.Context, _ *grpc.StreamDesc, name string, _ ...grpc.CallOption) (grpc.ClientStream, error) {
	localStream, ok := c.streams[name]
	if !ok {
		return nil, status.Errorf(codes.Unimplemented, "local stream %s not mounted", name)
	}
	pair := newStreamPair(ctx)
	go func() {
		err := localStream.desc.Handler(localStream.server, pair.server)
		pair.finish(err)
	}()
	return pair.client, nil
}

type streamPair struct {
	client *clientStream
	server *serverStream
}

func newStreamPair(ctx context.Context) *streamPair {
	clientToServer := make(chan streamMessage, 1)
	serverToClient := make(chan streamMessage, 1)
	state := &streamState{
		ctx:            ctx,
		clientToServer: clientToServer,
		serverToClient: serverToClient,
		done:           make(chan struct{}),
	}
	return &streamPair{
		client: &clientStream{state: state},
		server: &serverStream{state: state},
	}
}

func (p *streamPair) finish(err error) {
	p.client.state.finish(err)
}

type streamState struct {
	ctx            context.Context
	clientToServer chan streamMessage
	serverToClient chan streamMessage
	done           chan struct{}
	doneOnce       sync.Once
	closeSendOnce  sync.Once
	mu             sync.Mutex
	header         metadata.MD
	trailer        metadata.MD
	err            error
	sendClosed     bool
}

type streamMessage struct {
	msg proto.Message
	err error
}

func (s *streamState) finish(err error) {
	s.mu.Lock()
	s.err = err
	s.mu.Unlock()
	s.doneOnce.Do(func() {
		close(s.serverToClient)
		close(s.done)
	})
}

func (s *streamState) closeSend() {
	s.closeSendOnce.Do(func() {
		s.mu.Lock()
		s.sendClosed = true
		s.mu.Unlock()
		close(s.clientToServer)
	})
}

func (s *streamState) canSend() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return !s.sendClosed
}

func (s *streamState) statusErr() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.err == nil {
		return io.EOF
	}
	return s.err
}

type clientStream struct {
	state *streamState
}

func (s *clientStream) Header() (metadata.MD, error) {
	s.state.mu.Lock()
	defer s.state.mu.Unlock()
	return s.state.header.Copy(), nil
}

func (s *clientStream) Trailer() metadata.MD {
	s.state.mu.Lock()
	defer s.state.mu.Unlock()
	return s.state.trailer.Copy()
}

func (s *clientStream) CloseSend() error {
	s.state.closeSend()
	return nil
}

func (s *clientStream) Context() context.Context {
	return s.state.ctx
}

func (s *clientStream) SendMsg(m any) error {
	if !s.state.canSend() {
		return status.Error(codes.FailedPrecondition, "local stream send is closed")
	}
	msg, err := cloneMessage(m)
	if err != nil {
		return err
	}
	select {
	case <-s.state.done:
		return s.state.statusErr()
	case <-s.state.ctx.Done():
		return s.state.ctx.Err()
	case s.state.clientToServer <- streamMessage{msg: msg}:
		return nil
	}
}

func (s *clientStream) RecvMsg(m any) error {
	select {
	case item, ok := <-s.state.serverToClient:
		if !ok {
			return s.state.statusErr()
		}
		if item.err != nil {
			return item.err
		}
		return mergeMessage(m, item.msg)
	case <-s.state.ctx.Done():
		return s.state.ctx.Err()
	}
}

type serverStream struct {
	state *streamState
}

func (s *serverStream) SetHeader(md metadata.MD) error {
	s.state.mu.Lock()
	defer s.state.mu.Unlock()
	s.state.header = metadata.Join(s.state.header, md)
	return nil
}

func (s *serverStream) SendHeader(md metadata.MD) error {
	return s.SetHeader(md)
}

func (s *serverStream) SetTrailer(md metadata.MD) {
	s.state.mu.Lock()
	defer s.state.mu.Unlock()
	s.state.trailer = metadata.Join(s.state.trailer, md)
}

func (s *serverStream) Context() context.Context {
	return s.state.ctx
}

func (s *serverStream) SendMsg(m any) error {
	msg, err := cloneMessage(m)
	if err != nil {
		return err
	}
	select {
	case <-s.state.done:
		return s.state.statusErr()
	case <-s.state.ctx.Done():
		return s.state.ctx.Err()
	case s.state.serverToClient <- streamMessage{msg: msg}:
		return nil
	}
}

func (s *serverStream) RecvMsg(m any) error {
	select {
	case item, ok := <-s.state.clientToServer:
		if !ok {
			return io.EOF
		}
		if item.err != nil {
			return item.err
		}
		return mergeMessage(m, item.msg)
	case <-s.state.ctx.Done():
		return s.state.ctx.Err()
	}
}

func cloneMessage(m any) (proto.Message, error) {
	msg, ok := m.(proto.Message)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "local stream message is not proto message")
	}
	return proto.Clone(msg), nil
}

func mergeMessage(dst any, src proto.Message) error {
	msg, ok := dst.(proto.Message)
	if !ok {
		return status.Error(codes.InvalidArgument, "local stream target is not proto message")
	}
	proto.Reset(msg)
	proto.Merge(msg, src)
	return nil
}
