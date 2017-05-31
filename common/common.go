package common

import (
	"github.com/docker/docker/client"
)

const (
	BASE_DIR_DEFAULT string = "/tmp"
)

var (
	CLI      *client.Client
	BASE_DIR string = BASE_DIR_DEFAULT
)
