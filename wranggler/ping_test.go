package main_test

import (
	"context"
	"testing"
	"time"

	pb "worker/worker_proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// TestPing will test the ping handler for this worker.
func TestPing(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"bufnet",
		grpc.WithContextDialer(bufDialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewWorkerClient(conn)
	response, err := client.Ping(ctx, &pb.EmptyRequest{})
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}

	if response.GetStatus() != "pass" {
		t.Errorf("the response was not as expected: %v", response.Status)
	}
}
