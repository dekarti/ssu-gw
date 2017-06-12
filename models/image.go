package models

import (
	"github.com/dekarti/ssu-gw/util"
)

type (
	Image struct {
		Name string   `json:"name"`
		Tags []string `json:"tags"`
	}
)

func GetImages() ([]Image, error) {
	images, err := util.ListAvailableDockerImages()
	if err != nil {
		return nil, err
	}

	imageList := []Image{}
	for name, tags := range images {
		imageList = append(imageList, Image{
			Name: name,
			Tags: tags,
		})
	}
	return imageList, nil
}
