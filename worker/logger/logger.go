// logger provides utility functions for logging.
package logger

import (
	"context"
	"log"
	"strings"
	"time"
	pb "worker/logger_proto"

	"github.com/google/uuid"
)

type Logger struct {
	loggerClient pb.LoggerClient
	loggerToken  string
	environment  string
}

// New provides an initialization function for the logger package.
// It takes in the gRPC client of the logger, the token, and the environment
// of the application.
//
// Returns a pointer of type *Logger.
func New(client pb.LoggerClient, token string, env string) *Logger {
	return &Logger{
		loggerClient: client,
		loggerToken:  token,
		environment:  env,
	}
}

// Log provides a wrapper for the d.Logger.CreateLog function. More boilerplate, less work.
//
// If the requestID value is not provided (or it is an empty string), the function will
// automatically generate a new UUID for you.
// The body parameter might also be empty and would not cause any error.
//
// Usage:
//
//      Log("an error has occured!", logger.Level_ERROR.Enum(), "UUID", map[string]string{"key": "value"})
//
//      // or you can use it as a defer function to handle error cases
//      ...
//      if err != nil {
//          defer Log(err.Error(), logger.Level_ERROR.Enum(), "UUID", map[string]string{"key": "value"})
//          return err
//	    }
//
func (l *Logger) Log(message string, level *pb.Level, requestID string, body map[string]string) {
	// Handle panic case
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
		}
	}()

	loggerCtx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var LanguageGo = "Go"

	if requestID == "" {
		requestID = uuid.New().String()
	}

	_, requestErr := l.loggerClient.CreateLog(
		loggerCtx,
		&pb.LogRequest{
			AccessToken: l.loggerToken,
			Data: &pb.LogData{
				RequestId:   requestID,
				Application: "worker",
				Language:    &LanguageGo,
				Body:        body,
				Message:     message,
				Level:       level,
				Environment: l.GetLogEnvironment(),
			},
		},
	)
	if requestErr != nil {
		log.Printf(
			"An error has occured while trying to create a log to the logger service: %v\n\nTrying to send: %v",
			requestErr,
			message,
		)
	}
}

// GetLogEnvironment provides an utility function for passing the
// enum value of *logger.Environment for the log data.
func (l *Logger) GetLogEnvironment() *pb.Environment {
	switch strings.ToLower(l.environment) {
	case "development":
		return pb.Environment_DEVELOPMENT.Enum()
	case "production":
		return pb.Environment_PRODUCTION.Enum()
	case "testing":
		return pb.Environment_TESTING.Enum()
	case "staging":
		return pb.Environment_STAGING.Enum()
	default:
		return pb.Environment_UNSET.Enum()
	}
}
