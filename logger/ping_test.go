package main_test

import (
	"context"
	pb "logger/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
)

func TestPing(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)
	response, err := client.Ping(ctx, &pb.EmptyRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if response.GetStatus() != "pass" {
		t.Errorf("the response was not as expected: %v", response.Status)
	}
}
