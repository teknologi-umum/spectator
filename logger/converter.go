package main

import (
	pb "logger/proto"
	"time"
)

func (*Dependency) convertIntoLogData(l []*pb.LogData) []LogData {
	var data []LogData
	for _, d := range l {
		data = append(
			data,
			LogData{
				RequestID:   d.GetRequestId(),
				Application: d.GetApplication(),
				Message:     d.GetMessage(),
				Body:        d.GetBody(),
				Level:       d.GetLevel().String(),
				Environment: d.GetEnvironment().String(),
				Language:    d.GetLanguage(),
				Timestamp:   time.UnixMilli(d.GetTimestamp()),
			},
		)
	}

	return data
}

func (*Dependency) convertIntoProtoData(l []LogData) []*pb.LogData {
	var data []*pb.LogData
	for _, d := range l {
		var level pb.Level
		var environment pb.Environment
		var timestamp = d.Timestamp.UnixMilli()

		switch d.Level {
		case "INFO":
			level = pb.Level_INFO
		case "WARNING":
			level = pb.Level_WARNING
		case "ERROR":
			level = pb.Level_ERROR
		case "CRITICAL":
			level = pb.Level_CRITICAL
		default:
			level = pb.Level_DEBUG
		}

		switch d.Environment {
		case "PRODUCTION":
			environment = pb.Environment_PRODUCTION
		case "STAGING":
			environment = pb.Environment_STAGING
		case "DEVELOPMENT":
			environment = pb.Environment_DEVELOPMENT
		default:
			environment = pb.Environment_UNSET
		}

		data = append(data, &pb.LogData{
			RequestId:   d.RequestID,
			Application: d.Application,
			Message:     d.Message,
			Body:        d.Body,
			Level:       &level,
			Environment: &environment,
			Language:    &d.Language,
			Timestamp:   &timestamp,
		})
	}

	return data
}
