package handlers

import (
	"context"
	"fmt"
	//	"net/http"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	//	"github.com/docker/docker/api/types/strslice"
	"github.com/docker/docker/client"
	//	"github.com/labstack/echo"
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
