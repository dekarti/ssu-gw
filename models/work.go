package models

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	WORKING_DIR string = "/usr/src/myapp"
)

type Work struct {
	DockerClient *client.Client
	Task         *Task
	Path         string
	Input        string
	Output       string
	Command      []string
}

func (w Work) Validate() error {
	// Setup
	// if err := Unzip(w.Path); err != nil {
	//	return err
	//}

	// TODO:

	err := w.LaunchWork()
	if err != nil {
		return err
	}
	return nil

}

func (w Work) LaunchWork() error {
	containerName := fmt.Sprintf("work-%d-%d", w.Task.Number, time.Now().Unix())

	_, err := w.DockerClient.ContainerCreate(context.Background(),
		&container.Config{
			Cmd:        w.Command,
			Image:      "python:2.7-slim",
			WorkingDir: "/usr/src/myapp",
		},
		&container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:%s", w.Path, "/usr/src/myapp")},
		}, nil, containerName)

	if err != nil {
		return err
	}

	if err := w.DockerClient.ContainerStart(context.Background(), containerName, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.Mkdir(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
