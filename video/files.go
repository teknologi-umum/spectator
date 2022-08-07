package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/minio/minio-go/v7"
)

// acquireListOfFiles will get the list of files from the bucket
// and sort it by the name of the file based on the timestamp that
// is on the file's name. If the file's name has an invalid timestamp
// format, it will be skipped.
func (d *Dependency) acquireListOfFiles(ctx context.Context, sessionID string) ([]string, error) {
	defer func() {
		if e := recover(); e != nil {
			log.Print(e)
		}
	}()

	var files []string
	for f := range d.Bucket.ListObjects(ctx, sessionID, minio.ListObjectsOptions{}) {
		if !strings.HasSuffix(f.Key, ".webm") {
			continue
		}

		files = append(files, f.Key)
	}

	sort.SliceStable(files, func(i, j int) bool {
		// Split the file name into parts separated by underscore
		// and get the second part of the split.
		// This is the timestamp of the file.
		// We sort the files by timestamp.
		// Sample file name: 1659536466147_99881.webm

		file1 := strings.Split(strings.Replace(files[i], ".webm", "", 1), "_")
		file2 := strings.Split(strings.Replace(files[j], ".webm", "", 1), "_")

		timestamp1, err := strconv.ParseInt(file1[1], 10, 64)
		if err != nil {
			log.Printf("error parsing timestamp: %v", err)
		}

		timestamp2, err := strconv.ParseInt(file2[1], 10, 64)
		if err != nil {
			log.Printf("error parsing timestamp: %v", err)
		}

		// The earlier the timestamp, the earlier the file.
		return timestamp1 < timestamp2
	})

	return files, nil
}

// downloadFile with download ONE file from the bucket to the local
// container filesystem. It will return a file path to the file, that
// most likely will not be used, but hey, it's there.
func (d *Dependency) downloadFile(ctx context.Context, sessionID string, file string) (string, error) {
	filePath := path.Join(BaseDirectory, sessionID, file)

	object, err := d.Bucket.GetObject(ctx, sessionID, file, minio.GetObjectOptions{})
	if err != nil {
		return filePath, fmt.Errorf("getting the object: %w", err)
	}
	defer func() {
		err := object.Close()
		if err != nil {
			log.Printf("error closing object reader: %v", err)
		}
	}()

	objectBody, err := io.ReadAll(object)
	if err != nil {
		return filePath, fmt.Errorf("reading object: %w", err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		return filePath, fmt.Errorf("creating file: %w", err)
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Printf("error closing file: %v", err)
		}
	}()

	_, err = f.Write(objectBody)
	if err != nil {
		return filePath, fmt.Errorf("writing object into file: %w", err)
	}

	err = f.Sync()
	if err != nil {
		return filePath, fmt.Errorf("synchronizing file: %w", err)
	}

	return filePath, nil
}

// putListOfFilesToFile will put the list of files to a videos.txt file in the filesystem
// that can be used against a ffmpeg command to concatenate multiple videos.
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

// uploadCOmbinedFiles takes a filepath from the local (or container's filesystem)
// and uploads it to the bucket. It will return the bucket path for the resulting
// file.
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

// concatFiles will concatenate the list of files into one video using
// append to a single video based on the last offset of the file.
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

		_, err = f.Write(body)
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
