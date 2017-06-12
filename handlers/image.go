package handlers

import (
	"net/http"

	"github.com/dekarti/ssu-gw/models"
	"github.com/labstack/echo"
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

