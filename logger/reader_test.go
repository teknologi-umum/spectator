package main_test

import (
	"context"
	pb "logger/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestReadLog(t *testing.T) {
	t.Cleanup(cleanup)
	timeCurrent := time.Now().UnixMilli()
	payload := &pb.LogRequest{
		AccessToken: accessToken,
		Data: []*pb.LogData{
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
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	_, err = client.CreateLog(ctx, payload)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	resp, err := client.ReadLog(ctx, &pb.ReadLogRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if len(resp.Data) != 3 {
		t.Errorf("expected 3 logs, got %d", len(resp.Data))
	}
}

func TestReadLog_Empty(t *testing.T) {
	t.Cleanup(cleanup)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	resp, err := client.ReadLog(ctx, &pb.ReadLogRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if len(resp.Data) != 0 {
		t.Errorf("expected 0 logs, got %d", len(resp.Data))
	}
}
