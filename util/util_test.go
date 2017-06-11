package util

import (
	"context"
	"fmt"
	"os"
	"testing"
	"errors"

	"github.com/dekarti/ssu-gw/common"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
)

func TestMain(m *testing.M) {
	common.CLI, _ = client.NewEnvClient()
	os.Exit(m.Run())
}

func TestSearchDockerImage(t *testing.T) {
	javaImages, err := SearchDockerImage("java")
	t.Logf("%v", javaImages)
	assert.NoError(t, err)
	if len(javaImages) == 0 {
		err = errors.New("No images found.")
	}

	assert.NoError(t, err)

}

func TestPullDockerImage(t *testing.T) {
	reader, err := PullDockerImage("busybox:musl")
	assert.NoError(t, err)

	ioutil.ReadAll(reader)

	_, err = common.CLI.ImageRemove(context.Background(), "busybox:musl", types.ImageRemoveOptions{})
	assert.NoError(t, err)
}

func TestIsContainerSuccessfullyExited(t *testing.T) {
	testtable := []struct {
		In       string
		Expected bool
	}{
		{"true", true},
		{"false", false},
		{"true", true},
		//	{"sleep 100", false},
	}

	for i, testcase := range testtable {
		name := fmt.Sprintf("test-%d", i)
		err := RunContainer(common.CLI,
			&container.Config{
				Cmd:        []string{testcase.In},
				Image:      "python:2.7-slim",
				WorkingDir: "/usr/src/myapp",
			}, nil, name)
		assert.NoError(t, err)

		// TODO: Optimize it with goroutines
		err = IsContainerSuccessfullyExited(name, 4)
		flag := true
		if err != nil {
			flag = false
			assert.Equal(t, testcase.Expected, flag, "Error is %s", err)
		} else {
			assert.Equal(t, testcase.Expected, flag)
		}
		defer common.CLI.ContainerRemove(context.Background(), name, types.ContainerRemoveOptions{Force: true})
	}
}

func TestUnzip(t *testing.T) {
	// TODO
}
