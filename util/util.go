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

	"github.com/dekarti/ssu-gw/common"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func IsContainerSuccessfullyExited(name string, timeout int) error {
	containerInfo, err := common.CLI.ContainerInspect(context.Background(), name)
	for ; timeout > 0; timeout-- {
		time.Sleep(time.Second)
		if err != nil {
			return err
		}

		fmt.Printf("status: %s", containerInfo.State.Status)
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
