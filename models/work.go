package models

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/dekarti/ssu-gw/util"
	"github.com/docker/docker/api/types/container"
	"github.com/dekarti/ssu-gw/common"
)

type (
	Work struct {
		ContainerName string  `json:"containerName"`
		Task          *Task   `json:"task"`
		Path          string  `json:"path"`
		Input         string  `json:"input"`
		Output        string  `json:"output"`
		Config        *Config `json:"config"`
	}
)

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

	if err := util.RunContainer(common.CLI,
		&container.Config{
			Cmd:        w.Config.Command,
			Image:      fmt.Sprintf("%s:%s", w.Config.Language, w.Config.Version),
			WorkingDir: "/tmp",
		},
		&container.HostConfig{
			Binds: []string{fmt.Sprintf("%s:%s", w.Path, "/tmp")},
		}, containerName); err != nil {
		return err
	}

	w.ContainerName = containerName
	return nil
}
