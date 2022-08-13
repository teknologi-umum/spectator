package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"video/logger_proto"
	pb "video/video_proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const BaseDirectory = "/data"

func (d *Dependency) GetVideo(ctx context.Context, in *pb.VideoRequest) (*pb.VideoResponse, error) {
	log.Printf("Received video request: Session ID: %s", in.GetSessionId())

	err := d.Queue.Enqueue(&VideoJob{SessionId: in.GetSessionId()})
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "failed to enqueue video job: %v", err)
	}

	return &pb.VideoResponse{VideoUrl: ""}, nil
}

func (d *Dependency) executeVideoJob(sessionId string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	// Check if bucket exists
	exists, err := d.Bucket.BucketExists(ctx, sessionId)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("checking bucket existance: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("error: checking bucket existance: %v", err)
		return
	}

	if !exists {
		log.Printf("error: video data not found for session id: %s", sessionId)
		return
	}

	// Create the directory for current session ID
	err = os.MkdirAll(BaseDirectory+"/"+sessionId, 0755)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("creating directory: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("error: creating a directory: %v", err)
		return
	}

	files, err := d.acquireListOfFiles(ctx, sessionId)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("acquiring list of files: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("acquiring list of files: %v", err)
		return
	}

	defer func() {
		err = os.RemoveAll(BaseDirectory + "/" + sessionId)
		if err != nil {
			defer d.Logger.Log(
				fmt.Errorf("removing directory: %v", err).Error(),
				logger_proto.Level_ERROR.Enum(),
				"",
				map[string]string{
					"session_id": sessionId,
					"function":   "GetVideo",
				},
			)
			log.Printf("error: removing directory: %v", err)
		}
	}()

	for _, file := range files {
		_, err := d.downloadFile(ctx, sessionId, file)
		if err != nil {
			defer d.Logger.Log(
				fmt.Errorf("downloading file: %v", err).Error(),
				logger_proto.Level_ERROR.Enum(),
				"",
				map[string]string{
					"session_id": sessionId,
					"function":   "GetVideo",
				},
			)
			log.Printf("error: downloading file: %v", err)
			return
		}
	}

	outputCombinedWebmFile, err := d.concatFiles(sessionId, files)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("concatenating files: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("error: concatenating files: %v", err)
		return
	}

	uploadedWebmFilePath, err := d.uploadCombinedFile(ctx, sessionId, outputCombinedWebmFile)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("uploading combined file: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)

		log.Printf("error: uploading combined file: %v", err)
		return
	}

	outputMp4File := path.Join(BaseDirectory, sessionId, "combined.mp4")

	_, err = d.Ffmpeg.Convert(ctx, outputCombinedWebmFile, outputMp4File)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("converting webm file to mp4: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("error: converting webm to mp4: %v", err)
		return
	}

	uploadedMp4FilePath, err := d.uploadCombinedFile(ctx, sessionId, outputMp4File)
	if err != nil {
		defer d.Logger.Log(
			fmt.Errorf("uploading combined files to minio: %v", err).Error(),
			logger_proto.Level_ERROR.Enum(),
			"",
			map[string]string{
				"session_id": sessionId,
				"function":   "GetVideo",
			},
		)
		log.Printf("error: uploading combined files to minio %v", err)
		return
	}

	log.Printf("done: uploaded file for %s to %s and %s", sessionId, uploadedWebmFilePath, uploadedMp4FilePath)
}

type VideoJob struct {
	SessionId string
}

func VideoJobBuilder() interface{} {
	return &VideoJob{}
}
