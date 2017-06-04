package models

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dekarti/ssu-gw/util"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const (
	WORKING_DIR string = "/usr/src/myapp"
	TIMEOUT     int    = 5
)

type Work struct {
	DockerClient  *client.Client
	ContainerName string
	Task          *Task
	Path          string
	Input         string
	Output        string
	Command       []string
}

func (w *Work) WriteInput(s string) error {
	f, err := os.Create(fmt.Sprintf("%s/input.txt", w.Path))
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString(s)
	writer.Flush()
	return err
}

func (w *Work) ReadInput() (string, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/input.txt", w.Path))
	return string(b), err
}

func (w *Work) ReadOutput() (string, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/output.txt", w.Path))
	return string(b), err
}

func (w *Work) Validate() error {
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

func (w *Work) LaunchWork() error {
	containerName := fmt.Sprintf("work-%d-%d", w.Task.Id, time.Now().Unix())

	if err := util.RunContainer(w.DockerClient,
		&container.Config{
			Cmd:        w.Command,
			Image:      "python:2.7-slim",
			WorkingDir: "/usr/src/myapp",
		},
		&container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:%s", w.Path, "/usr/src/myapp")},
		}, containerName); err != nil {
		return err
	}

	w.ContainerName = containerName
	return nil
}
