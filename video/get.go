package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"time"

	"video/logger_proto"
	pb "video/video_proto"

	"github.com/minio/minio-go/v7"
)

const BaseDirectory = "/data"

func (d *Dependency) GetVideo(ctx context.Context, in *pb.VideoRequest) (*pb.VideoResponse, error) {
	log.Printf("Received video request: Session ID: %s", in.GetSessionId())

	go func(sessionId string) {
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
		err = os.MkdirAll(BaseDirectory+"/"+in.SessionId, 0755)
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

		for _, file := range files {
			err = d.downloadFile(ctx, sessionId, file)
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

		uploadedFilePath, err := d.uploadCombinedFile(ctx, sessionId, outputMp4File)
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
			return
		}

		log.Printf("done: uploaded file for %s to %s", sessionId, uploadedFilePath)
	}(in.GetSessionId())

	return &pb.VideoResponse{VideoUrl: ""}, nil
}

func (d *Dependency) acquireListOfFiles(ctx context.Context, sessionID string) ([]string, error) {
	defer func() {
		if e := recover(); e != nil {
			log.Print(e)
		}
	}()

	var files []string
	for f := range d.Bucket.ListObjects(ctx, sessionID, minio.ListObjectsOptions{}) {
		files = append(files, f.Key)
	}

	return files, nil
}

func (d *Dependency) downloadFile(ctx context.Context, sessionID string, file string) error {
	filePath := path.Join(BaseDirectory, sessionID, file)

	object, err := d.Bucket.GetObject(ctx, sessionID, file, minio.GetObjectOptions{})
	if err != nil {
		return fmt.Errorf("getting the object: %w", err)
	}
	defer func() {
		err := object.Close()
		if err != nil {
			log.Printf("error closing object reader: %v", err)
		}
	}()

	objectBody, err := io.ReadAll(object)
	if err != nil {
		return fmt.Errorf("reading object: %w", err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	_, err = f.Write(objectBody)
	if err != nil {
		return fmt.Errorf("writing object into file: %w", err)
	}

	err = f.Sync()
	if err != nil {
		return fmt.Errorf("synchronizing file: %w", err)
	}

	return nil
}

func (d *Dependency) putListOfFilesToFile(sessionID string, files []string) (string, error) {
	filePath := path.Join(BaseDirectory, sessionID, "videos.txt")
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	for _, file := range files {
		_, err := f.WriteString("file '" + file + "'\n")
		if err != nil {
			return "", fmt.Errorf("writing buffer to file: %w", err)
		}
	}

	err = f.Sync()
	if err != nil {
		return "", fmt.Errorf("synchronizing file: %w", err)
	}

	return filePath, nil
}

func (d *Dependency) uploadCombinedFile(ctx context.Context, sessionID string, filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("opening file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	fileStat, err := f.Stat()
	if err != nil {
		return "", fmt.Errorf("getting file stat: %w", err)
	}

	resultingFile := sessionID + "_" + "video.mp4"
	_, err = d.Bucket.PutObject(ctx, "public", resultingFile, f, fileStat.Size(), minio.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("putting object: %w", err)
	}

	return path.Join("public", resultingFile), nil
}

func (d *Dependency) concatFiles(id string, files []string) (string, error) {
	outputFilePath := path.Join(BaseDirectory, id, "combined.webm")

	f, err := os.Create(outputFilePath)
	if err != nil {
		return "", fmt.Errorf("creating file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	for _, file := range files {
		r, err := os.Open(path.Join(BaseDirectory, id, file))
		if err != nil {
			return "", fmt.Errorf("opening file: %w", err)
		}
		defer func() {
			err := r.Close()
			if err != nil {
				log.Printf("error closing file: %v", err)
			}
		}()

		body, err := io.ReadAll(r)
		if err != nil {
			return "", fmt.Errorf("reading file: %w", err)
		}

		fileInfo, err := f.Stat()
		if err != nil {
			return "", fmt.Errorf("reading file status: %w", err)
		}

		_, err = f.WriteAt(body, fileInfo.Size())
		if err != nil {
			return "", fmt.Errorf("writing buffer to file: %w", err)
		}
	}

	err = f.Sync()
	if err != nil {
		return "", fmt.Errorf("synchronizing file: %w", err)
	}

	return outputFilePath, nil
}
