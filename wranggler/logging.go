package main

// logging.go provides utility functions for logging.

import (
	"context"
	"log"
	"strings"
	"time"
	"worker/logger"

	"github.com/google/uuid"
)

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
func (d *Dependency) Log(message string, level *logger.Level, requestID string, body map[string]string) {
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

	_, requestErr := d.Logger.CreateLog(
		loggerCtx,
		&logger.LogRequest{
			AccessToken: d.LoggerToken,
			Data: &logger.LogData{
				RequestId:   requestID,
				Application: "worker",
				Language:    &LanguageGo,
				Body:        body,
				Message:     message,
				Level:       level,
				Environment: d.GetLogEnvironment(),
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
func (d *Dependency) GetLogEnvironment() *logger.Environment {
	switch strings.ToLower(d.Environment) {
	case "development":
		return logger.Environment_DEVELOPMENT.Enum()
	case "production":
		return logger.Environment_PRODUCTION.Enum()
	case "testing":
		return logger.Environment_TESTING.Enum()
	case "staging":
		return logger.Environment_STAGING.Enum()
	default:
		return logger.Environment_UNSET.Enum()
	}
}
