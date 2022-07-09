package main

import (
	"context"
	"log"
	"strings"
	"time"
	"video/logger_proto"

	"github.com/google/uuid"
)

type Logger struct {
	client      logger_proto.LoggerClient
	token       string
	environment string
	noop        bool
}

func NewLogger(client logger_proto.LoggerClient, token string, environment string) *Logger {
	if strings.ToLower(environment) != "production" {
		return &Logger{noop: true}
	}

	return &Logger{
		client: client,
		token:  token,
		noop:   false,
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
func (l *Logger) Log(message string, level *logger_proto.Level, requestID string, body map[string]string) {
	if l.noop {
		log.Println(message)
		return
	}

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

	_, requestErr := l.client.CreateLog(
		loggerCtx,
		&logger_proto.LogRequest{
			AccessToken: l.token,
			Data: &logger_proto.LogData{
				RequestId:   requestID,
				Application: "video",
				Language:    &LanguageGo,
				Body:        body,
				Message:     message,
				Level:       level,
				Environment: l.getLogEnvironment(),
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

// getLogEnvironment provides an utility function for passing the
// enum value of *logger.Environment for the log data.
func (l *Logger) getLogEnvironment() *logger_proto.Environment {
	switch strings.ToLower(l.environment) {
	case "development":
		return logger_proto.Environment_DEVELOPMENT.Enum()
	case "production":
		return logger_proto.Environment_PRODUCTION.Enum()
	case "testing":
		return logger_proto.Environment_TESTING.Enum()
	case "staging":
		return logger_proto.Environment_STAGING.Enum()
	default:
		return logger_proto.Environment_UNSET.Enum()
	}
}
