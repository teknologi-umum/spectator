package main_test

import (
	"context"
	logger "logger"
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

func TestValidatePayload(t *testing.T) {
	deps := logger.Dependency{
		DB:          db,
		Org:         influxOrg,
		AccessToken: accessToken,
	}

	t.Run("empty", func(t *testing.T) {
		p := &pb.LogRequest{}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "access token must be provided" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})

	t.Run("missing", func(t *testing.T) {
		p := &pb.LogRequest{AccessToken: accessToken}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "proper request_id, application, message must be provided" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})

	t.Run("commas", func(t *testing.T) {
		p := &pb.LogRequest{
			AccessToken: accessToken,
			Data: []*pb.LogData{{
				RequestId:   "bla,bla",
				Application: "asd,asd",
				Message:     "hello there",
			}},
		}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "proper request_id, application must be provided" {
			t.Errorf("expecting an error, nistead got: %v", err)
		}
	})
}

func TestSingleCreateRead(t *testing.T) {
	// TODO

	_ = []logger.LogPayload{
		{
			AccessToken: accessToken,
			Data: []logger.LogData{
				{
					RequestID:   "a1",
					Application: "core",
					Message:     "A quick brown fox jumps over the lazy dog",
					Level:       "info",
					Environment: "production",
					Language:    "C#",
					Timestamp:   time.Now(),
				},
			},
		},
		{
			AccessToken: accessToken,
			Data: []logger.LogData{{
				RequestID:   "a1",
				Application: "worker",
				Message:     "Oh no, something went wrong",
				Level:       "error",
				Environment: "production",
				Language:    "Javascript",
				Body: map[string]string{
					"stack_trace": "file.js:70 anotherfile.js:30",
					"why":         "I don't know",
				},
			}},
		},
		{
			AccessToken: accessToken,
			Data: []logger.LogData{{
				RequestID:   "b2",
				Application: "core",
				Message:     "Well, hello there. General Kenobi.",
			}},
		},
		{
			AccessToken: accessToken,
			Data: []logger.LogData{{
				RequestID:   "c3",
				Application: "worker",
				Message:     "This happened in the past",
				Timestamp:   time.Now().Add(time.Hour * 6 * -1),
			}},
		},
	}
}

func TestBulkCreate(t *testing.T) {
	// TODO
}
