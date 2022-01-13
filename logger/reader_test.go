package main_test

import (
	"context"
	pb "logger/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestReadLog(t *testing.T) {
	t.Cleanup(cleanup)
	timeCurrent := time.Now().UnixMilli()
	payload := []*pb.LogData{
		{
			RequestId:   "a1",
			Application: "core",
			Message:     "A quick brown fox jumps over the lazy dog",
			Level:       pb.Level_DEBUG.Enum(),
			Environment: pb.Environment_PRODUCTION.Enum(),
			Timestamp:   &timeCurrent,
		},
		{
			RequestId:   "b2",
			Application: "core",
			Message:     "Well, hello there. General Kenobi.",
			Level:       pb.Level_DEBUG.Enum(),
			Environment: pb.Environment_STAGING.Enum(),
			Timestamp:   &timeCurrent,
		},
		{
			RequestId:   "c3",
			Application: "worker",
			Message:     "Lorem ipsum dolor sit amet",
			Level:       pb.Level_DEBUG.Enum(),
			Environment: pb.Environment_DEVELOPMENT.Enum(),
			Timestamp:   &timeCurrent,
		},
		{
			RequestId:   "d4",
			Application: "piston",
			Message:     "someone executed this",
			Level:       pb.Level_DEBUG.Enum(),
			Environment: pb.Environment_DEVELOPMENT.Enum(),
			Body: map[string]string{
				"foo": "bar",
				"baz": "qux",
			},
			Timestamp: &timeCurrent,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	for _, logData := range payload {
		_, err = client.CreateLog(ctx, &pb.LogRequest{
			AccessToken: accessToken,
			Data:        logData,
		})
		if err != nil {
			t.Errorf("an error was thrown: %v", err)
		}
	}

	resp, err := client.ReadLog(ctx, &pb.ReadLogRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if len(resp.GetData()) != 4 {
		t.Errorf("expected 4 logs, got %d", len(resp.GetData()))
	}
}

func TestReadLog_Empty(t *testing.T) {
	t.Cleanup(cleanup)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	resp, err := client.ReadLog(ctx, &pb.ReadLogRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if len(resp.GetData()) != 0 {
		t.Errorf("expected 0 logs, got %d", len(resp.GetData()))
	}
}

func TestReadLog_Query(t *testing.T) {
	t.Skip("need to fix something up on the reader.go")
	t.Cleanup(cleanup)

	payload := []*pb.LogData{
		{
			RequestId:   "a1",
			Application: "core",
			Message:     "A quick brown fox jumps over the lazy dog",
			Level:       pb.Level_ERROR.Enum(),
			Environment: pb.Environment_PRODUCTION.Enum(),
		},
		{
			RequestId:   "a1",
			Application: "core",
			Message:     "Well, hello there. General Kenobi.",
			Level:       pb.Level_ERROR.Enum(),
			Environment: pb.Environment_PRODUCTION.Enum(),
		},
		{
			RequestId:   "a1",
			Application: "worker",
			Message:     "Lorem ipsum dolor sit amet",
			Level:       pb.Level_ERROR.Enum(),
			Environment: pb.Environment_PRODUCTION.Enum(),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	for _, logData := range payload {
		_, err = client.CreateLog(ctx, &pb.LogRequest{
			AccessToken: accessToken,
			Data:        logData,
		})
		if err != nil {
			t.Errorf("an error was thrown: %v", err)
		}

	}

	a1 := "a1"
	core := "core"
	resp, err := client.ReadLog(
		ctx,
		&pb.ReadLogRequest{
			Level:       pb.Level_ERROR.Enum(),
			Application: &core,
			RequestId:   &a1,
		},
	)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if len(resp.GetData()) != 2 {
		t.Errorf("expected 2 log, got %d", len(resp.GetData()))
	}

}
