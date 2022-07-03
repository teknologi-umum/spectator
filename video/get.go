package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"

	pb "video/video_proto"

	"github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const BaseDirectory = "/data"

func (d *Dependency) GetVideo(ctx context.Context, in *pb.VideoRequest) (*pb.VideoResponse, error) {
	// Check if bucket exists
	exists, err := d.Bucket.BucketExists(ctx, in.SessionId)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "checking bucket existance: %v", err)
	}

	if !exists {
		return &pb.VideoResponse{}, status.Errorf(codes.NotFound, "video data not found for session id: %s", in.SessionId)
	}

	// Create the directory for current session ID
	err = os.MkdirAll(BaseDirectory+"/"+in.SessionId, 0755)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "creating a directory: %v", err)
	}

	files, err := d.acquireListOfFiles(ctx, in.SessionId)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "acquiring list of files: %v", err)
	}

	for _, file := range files {
		err = d.downloadFile(ctx, in.SessionId, file)
		if err != nil {
			return &pb.VideoResponse{}, status.Errorf(codes.Internal, "downloading file: %v", err)
		}
	}

	listFilePath, err := d.putListOfFilesToFile(in.SessionId, files)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "generating file lists: %v", err)
	}

	outputCombinedWebmFile := path.Join(BaseDirectory, in.SessionId, "combined.webm")

	_, err = d.Ffmpeg.Concat(ctx, listFilePath, outputCombinedWebmFile)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "combining files with ffmpeg: %v", err)
	}

	outputMp4File := path.Join(BaseDirectory, in.SessionId, "combined.mp4")

	_, err = d.Ffmpeg.Convert(ctx, outputCombinedWebmFile, outputMp4File)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "converting webm to mp4: %v", err)
	}

	uploadedFilePath, err := d.uploadCombinedFile(ctx, in.SessionId, outputMp4File)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "uploading combined files to minio %v", err)
	}

	err = os.RemoveAll(BaseDirectory + "/" + in.SessionId)
	if err != nil {
		return &pb.VideoResponse{}, status.Errorf(codes.Internal, "removing directory: %v", err)
	}

	return &pb.VideoResponse{VideoUrl: uploadedFilePath}, nil
}

func (d *Dependency) acquireListOfFiles(ctx context.Context, sessionID string) ([]string, error) {
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
