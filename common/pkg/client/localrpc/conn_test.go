package localrpc

import (
	"context"
	"errors"
	"io"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type testStreamService struct{}

func TestClientStream(t *testing.T) {
	conn := NewConn()
	conn.RegisterService(&grpc.ServiceDesc{
		ServiceName: "test.StreamService",
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Sum",
				ClientStreams: true,
				Handler: func(_ any, stream grpc.ServerStream) error {
					var sum int64
					for {
						req := new(wrapperspb.Int64Value)
						if err := stream.RecvMsg(req); errors.Is(err, io.EOF) {
							return stream.SendMsg(wrapperspb.Int64(sum))
						} else if err != nil {
							return err
						}
						sum += req.GetValue()
					}
				},
			},
		},
	}, &testStreamService{})

	stream, err := conn.NewStream(context.Background(), nil, "/test.StreamService/Sum")
	if err != nil {
		t.Fatal(err)
	}
	for _, value := range []int64{1, 2, 3} {
		if err := stream.SendMsg(wrapperspb.Int64(value)); err != nil {
			t.Fatal(err)
		}
	}
	if err := stream.CloseSend(); err != nil {
		t.Fatal(err)
	}

	resp := new(wrapperspb.Int64Value)
	if err := stream.RecvMsg(resp); err != nil {
		t.Fatal(err)
	}
	if resp.GetValue() != 6 {
		t.Fatalf("sum = %d, want 6", resp.GetValue())
	}
	if err := stream.RecvMsg(new(wrapperspb.Int64Value)); !errors.Is(err, io.EOF) {
		t.Fatalf("final recv error = %v, want EOF", err)
	}
}

func TestServerStream(t *testing.T) {
	conn := NewConn()
	conn.RegisterService(&grpc.ServiceDesc{
		ServiceName: "test.StreamService",
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "Range",
				ServerStreams: true,
				Handler: func(_ any, stream grpc.ServerStream) error {
					req := new(wrapperspb.Int64Value)
					if err := stream.RecvMsg(req); err != nil {
						return err
					}
					for i := int64(0); i < req.GetValue(); i++ {
						if err := stream.SendMsg(wrapperspb.Int64(i)); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}, &testStreamService{})

	stream, err := conn.NewStream(context.Background(), nil, "/test.StreamService/Range")
	if err != nil {
		t.Fatal(err)
	}
	if err := stream.SendMsg(wrapperspb.Int64(3)); err != nil {
		t.Fatal(err)
	}
	if err := stream.CloseSend(); err != nil {
		t.Fatal(err)
	}
	for i := int64(0); i < 3; i++ {
		resp := new(wrapperspb.Int64Value)
		if err := stream.RecvMsg(resp); err != nil {
			t.Fatal(err)
		}
		if resp.GetValue() != i {
			t.Fatalf("value = %d, want %d", resp.GetValue(), i)
		}
	}
	if err := stream.RecvMsg(new(wrapperspb.Int64Value)); !errors.Is(err, io.EOF) {
		t.Fatalf("final recv error = %v, want EOF", err)
	}
}
