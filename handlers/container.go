package handlers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dekarti/ssu-gw/common"
	"github.com/dekarti/ssu-gw/models"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	//	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	"github.com/labstack/echo"
)

const (
	VOLUME_PATH string = "/Users/aarutyunyan/Projects/study/ssu-fl/tasks"
	WORKING_DIR string = "/usr/src/myapp"
)

//func HelloHandler(c echo.Context) error {
//	response := HelloResponse{
//		Message: "Hello",
//	}
//	return c.JSON(http.StatusOK, response)
//}

func LaunchWorkHandler(c echo.Context) error {
	work := &models.Work{
		DockerClient: common.CLI,
		Task:         models.Task1,
		Path:         "/Users/aarutyunyan/Desktop/ssu-fl/tasks",
		Command:      []string{"python", "exp_nfa.py"},
	}

	if err := work.WriteInput(models.Task1.DefaultInput); err != nil {
		return err
	}

	if err := work.LaunchWork(); err != nil {
		return err
	}

	if output, err := work.ReadOutput(); err != nil {
		return err
	} else if output != models.Task1.ExpectedOutput {
		return c.JSON(http.StatusBadRequest, struct {
			Message  string
			Input    string
			Output   string
			Expected string
		}{
			Message:  "doesn't match",
			Input:    models.Task1.DefaultInput,
			Output:   output,
			Expected: models.Task1.ExpectedOutput,
		})
	} else {
		return c.JSON(http.StatusOK, struct {
			Message  string
			Input    string
			Output   string
			Expected string
		}{
			Message:  "ok",
			Input:    models.Task1.DefaultInput,
			Output:   output,
			Expected: models.Task1.ExpectedOutput,
		})
	}
}

func RunContainer(cli *client.Client) (*container.ContainerCreateCreatedBody, error) {
	response, err := cli.ContainerCreate(context.Background(),
		&container.Config{
			Cmd:        []string{"python", "exp_nfa.py"},
			Image:      "python:2.7-slim",
			WorkingDir: WORKING_DIR,
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: VOLUME_PATH,
					Target: WORKING_DIR,
				},
			},
		}, nil, fmt.Sprintf("my-container-%d", time.Now().Unix()))

	if err != nil {
		return nil, err
	}
	return &response, nil
}
