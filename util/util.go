package util

import (
	"archive/zip"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/dekarti/ssu-gw/common"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/filters"
)

func PullDockerImage(name string) (io.ReadCloser, error) {
	if reader, err := common.CLI.ImagePull(context.Background(), name, types.ImagePullOptions{}); err != nil {
		return nil, err
	} else {
		return reader, nil
	}
}

func SearchDockerImage(name string) ([]string, error) {
	images, err := common.CLI.ImageSearch(context.Background(), name, types.ImageSearchOptions{
		Limit: 10,
	})
	if err != nil {
		return nil, err
	}

	result := []string{}
	for _, image := range images {
		result = append(result, image.Name)
	}
	return result, nil
}

func SearchDockerImageTags(name string) ([]string, error) {
	absoluteName := fmt.Sprintf("library/%s", name)
	if strings.Contains(name, "/") {
		absoluteName = name
	}

	next := fmt.Sprintf("https://registry.hub.docker.com/v2/repositories/%s/tags/", absoluteName)
	result := []string{}
	for next != "" {
		resp, err := http.Get(next)
		if err != nil {
			return nil, err
		}

		responseData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		tags := struct {
			Next    string
			Results []struct {
				Name string `json:"name"`
			} `json:"results"`
		}{}

		if err := json.Unmarshal(responseData, &tags); err != nil {
			return nil, err
		}

		for _, tag := range tags.Results {
			result = append(result, tag.Name)
		}
		next = tags.Next
	}
	return result, nil
}

func ListAvailableDockerImages() (map[string][]string, error) {
	filterArgs := filters.NewArgs()
	filterArgs.Add("dangling", "false")
	images, err := common.CLI.ImageList(context.Background(), types.ImageListOptions{
		Filters: filterArgs,
	})
	if err != nil {
		return nil, err
	}

	imageMap := map[string][]string{}
	for _, image := range images {
		for _, imageName := range image.RepoTags {
			tmp := strings.Split(imageName, ":")
			imagePrefix := tmp[0]
			imageSuffix := tmp[1:]
			if _, ok := imageMap[imagePrefix]; ok {
				imageMap[imagePrefix] = append(imageMap[imagePrefix], imageSuffix...)
			} else {
				imageMap[imagePrefix] = imageSuffix
			}
		}
	}

	return imageMap, nil
}

func IsContainerSuccessfullyExited(name string, timeout int) error {
	containerInfo, err := common.CLI.ContainerInspect(context.Background(), name)
	for ; timeout > 0; timeout-- {
		time.Sleep(time.Second)
		if err != nil {
			return err
		}

		if containerInfo.State.Status == "exited" {
			if containerInfo.State.ExitCode != 0 {
				return errors.New(fmt.Sprintf("Process exited with code %d", containerInfo.State.ExitCode))
			}
			return nil
		}
		var err error
		containerInfo, err = common.CLI.ContainerInspect(context.Background(), name)
		if err != nil {
			return err
		}
	}
	return errors.New(fmt.Sprintf("Timeout error. Container status is %s", containerInfo.State.Status))
}

func RunContainer(cli *client.Client, cfg *container.Config, hcfg *container.HostConfig, name string) error {
	if _, err := cli.ContainerCreate(context.Background(), cfg, hcfg, nil, name); err != nil {
		return err
	}
	if err := cli.ContainerStart(context.Background(), name, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	switch err := os.Mkdir(target, 0755); {
	case os.IsExist(err):
		break
	case err != nil:
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
