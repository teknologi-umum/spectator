package main_test

import (
	"context"
	logger "logger"
	pb "logger/proto"
	"testing"
	"time"

	"google.golang.org/grpc"
)

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
		if err == nil || err.Error() != "proper data must be provided" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})

	t.Run("empty", func(t *testing.T) {
		p := &pb.LogRequest{
			AccessToken: accessToken,
			Data: []*pb.LogData{
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
			},
		}
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

	t.Run("five_limit", func(t *testing.T) {
		p := &pb.LogRequest{
			AccessToken: accessToken,
			Data: []*pb.LogData{
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
				{
					RequestId:   "",
					Application: "",
					Message:     "",
				},
			},
		}
		err := deps.ValidatePayload(p)
		if err == nil || err.Error() != "log data is more than five, maximum of five" {
			t.Errorf("expecting an error, instead got: %v", err)
		}
	})
}

func TestSingleCreate(t *testing.T) {
	t.Cleanup(cleanup)
	timePast := time.Now().Add(time.Hour * 6 * -1).UnixMilli()
	timeCurrent := time.Now().UnixMilli()
	timeZero := time.Unix(0, 0).UnixMilli()
	seesharp := "C#"
	javascreep := "Javascript"
	payloads := []pb.LogRequest{
		{
			AccessToken: accessToken,
			Data: []*pb.LogData{{
				RequestId:   "a1",
				Application: "core",
				Message:     "A quick brown fox jumps over the lazy dog",
				Level:       pb.Level_INFO.Enum(),
				Environment: pb.Environment_PRODUCTION.Enum(),
				Language:    &seesharp,
				Timestamp:   &timeCurrent,
			}},
		},
		{
			AccessToken: accessToken,
			Data: []*pb.LogData{{
				RequestId:   "a1",
				Application: "worker",
				Message:     "Oh no, something went wrong",
				Level:       pb.Level_ERROR.Enum(),
				Environment: pb.Environment_PRODUCTION.Enum(),
				Language:    &javascreep,
				Body: map[string]string{
					"stack_trace": "file.js:70 anotherfile.js:30",
					"why":         "I don't know",
				},
			}},
		},
		{
			AccessToken: accessToken,
			Data: []*pb.LogData{{
				RequestId:   "b2",
				Application: "core",
				Message:     "Well, hello there. General Kenobi.",
				Timestamp:   &timeZero,
			}},
		},
		{
			AccessToken: accessToken,
			Data: []*pb.LogData{{
				RequestId:   "c3",
				Application: "worker",
				Message:     "This happened in the past",
				Timestamp:   &timePast,
			}},
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Errorf("an error was thrown: %v", err)
	}
	defer conn.Close()

	client := pb.NewLoggerClient(conn)

	for i := 0; i < len(payloads); i++ {
		_, err = client.CreateLog(ctx, &payloads[i])
		if err != nil {
			t.Errorf("[#%d] an error was thrown: %v", i, err)
		}
	}
}

func TestBulkCreate(t *testing.T) {
	t.Cleanup(cleanup)
	timeCurrent := time.Now().UnixMilli()
	payload := &pb.LogRequest{
		AccessToken: accessToken,
		Data: []*pb.LogData{
			{
				RequestId:   "a1",
				Application: "core",
				Message:     "A quick brown fox jumps over the lazy dog",
				Level:       pb.Level_INFO.Enum(),
				Environment: pb.Environment_PRODUCTION.Enum(),
			},
			{
				RequestId:   "b2",
				Application: "core",
				Message:     "Well, hello there. General Kenobi.",
				Level:       pb.Level_WARNING.Enum(),
				Environment: pb.Environment_STAGING.Enum(),
			},
			{
				RequestId:   "c3",
				Application: "worker",
				Message:     "Lorem ipsum dolor sit amet",
				Level:       pb.Level_ERROR.Enum(),
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
}
