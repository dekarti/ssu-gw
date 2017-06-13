package handlers

import (
	"net/http"
	"io/ioutil"

	"github.com/dekarti/ssu-gw/models"
	"github.com/dekarti/ssu-gw/util"
	"github.com/labstack/echo"
	"fmt"
)

//GET  /images
func GetImagesHandler(c echo.Context) error {
	images, err := models.GetImages()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Images []models.Image `json:"images"`
	}{
		Images: images,
	})
}

//GET /images/search?name=
func SearchImageHandler(c echo.Context) error {
	official := c.QueryParam("official")
	imageName := c.QueryParam("name")
	images, err := util.SearchDockerImage(imageName, official == "1")
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Images []string `json:"images"`
	}{
		Images: images,
	})
}

//GET /images/:name/tags
func ListImageTagsHandler(c echo.Context) error {
	var imageName string
	if c.Param("namesuffix") == "" {
		imageName = c.Param("name")
	} else {
		imageName = fmt.Sprintf("%s/%s", c.Param("nameprefix"), c.Param("namesuffix"))
	}
	tags, err := util.ListDockerImageTags(imageName)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, struct {
		Tags []string `json:"tags"`
	}{
		Tags: tags,
	})
}

//GET /images/pull?name=
func PullImageHandler(c echo.Context) error {
	imageName := c.QueryParam("name")
	if reader, err := util.PullDockerImage(imageName); err != nil {
		return err
	} else {
		ioutil.ReadAll(reader)
	}

	return c.String(http.StatusOK, imageName)
}
