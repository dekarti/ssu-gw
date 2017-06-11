package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	_ "strconv"
	_ "strings"
	"time"

	"github.com/dekarti/ssu-gw/common"
	"github.com/dekarti/ssu-gw/models"
	"github.com/dekarti/ssu-gw/util"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// POST /upload
func UploadHandler(c echo.Context) error {
	form, err := c.MultipartForm()
	for k, _ := range form.File {
		log.Debug(k)
	}
	file, err := c.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	src, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// Destination
	// TODO: Generate work dir depending on work parameters
	//	 Now it uses timestamp.
	workDir := fmt.Sprintf("work_%d", time.Now().Unix())
	dstDir := fmt.Sprintf("%s/%s", common.BASE_DIR, workDir)
	os.Mkdir(dstDir, 0755)
	dst, err := os.Create(fmt.Sprintf("%s/%s", dstDir, file.Filename))
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	// Copy

	if _, err = io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

	if err = util.Unzip(dst.Name(), dstDir); err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, struct {
		WorkPath string `json:"path"`
	}{
		dstDir,
	})
}

// POST /work/launch
func LaunchWorkHandler(c echo.Context) error {
	//uaskNumber, _ := strconv.Atoi(c.Param("id"))
	work := &models.Work{}
	if err := c.Bind(work); err != nil {
		return err
	}

	log.Debugf("binded")

	if work.Input == "" {
		work.Input = work.Task.DefaultInput
	}

	if err := work.WriteInput(work.Input); err != nil {
		return err
	}

	if err := work.LaunchWork(); err != nil {
		return err
	}

	if err := util.IsContainerSuccessfullyExited(work.ContainerName, 5); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if output, err := work.ReadOutput(); err != nil {
		return err
	} else {
		work.Output = output
		/*
		if work.Expected != "" {
			if output == testUnit.Expected {
				testUnit.Result = "Output matches expected value"
				return c.JSON(http.StatusOK, testUnit)
			} else {
				testUnit.Result = "Output doesn't match expected value"
				return c.JSON(http.StatusBadRequest, testUnit)
			}
		}*/
		return c.JSON(http.StatusOK, work)
	}
}
