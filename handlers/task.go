package handlers

import (
	"net/http"
	"time"
	"math/rand"

	"github.com/dekarti/ssu-gw/models"
	"github.com/labstack/echo"
)


// GET 	/tasks
func GetTasksHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Tasks []*models.Task `json:"tasks"`
	}{
		models.Tasks,
	})
}

// POST /tasks
func CreateTaskHandler(c echo.Context) error {
	newTask := &models.Task{}

	if err := c.Bind(newTask); err != nil {
		return err
	}

	rand.Seed(time.Now().UTC().UnixNano())
	newTask.Id = rand.Intn(100000)

	models.Tasks = append(models.Tasks, newTask)

	return c.JSON(http.StatusCreated, newTask)
}
