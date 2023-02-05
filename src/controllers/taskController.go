package controllers

import (
	"gogolook-interview/models"
	"gogolook-interview/repositorys"
	"gogolook-interview/requests"
	"gogolook-interview/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status int8   `json:"status"`
}

func GetTasks(c *gin.Context) {
	var tasks []Task

	/* get task from DB */
	taskModels, err := repositorys.GetTasks()

	if err != nil {
		response(c, http.StatusBadRequest, err)
	}

	for _, taskModel := range taskModels {
		tasks = append(tasks, transform(taskModel))
	}

	response(c, http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
	var req requests.CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	/* save task to DB */
	newTaskModel, err := repositorys.CreateTask(req)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())
	}

	response(c, http.StatusOK, transform(newTaskModel))
}

func UpdateTask(c *gin.Context) {
	var req requests.UpdateTaskRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	taskModel, err := repositorys.UpdateTask(req)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())

		return
	}

	response(c, http.StatusCreated, transform(taskModel))
}
func DeleteTask(c *gin.Context) {
	var req requests.DeleteTaskRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := repositorys.DeleteTask(req)
	if err != nil {
		response(c, http.StatusBadRequest, err.Error())

		return
	}

	response(c, http.StatusOK, nil)
}

func response(c *gin.Context, status int, response interface{}) {
	c.JSON(status, gin.H{"result": response})
}

func transform(model models.Task) Task {
	return Task{
		ID:     model.ID,
		Name:   model.Name,
		Status: utils.BoolToInt8(model.Status),
	}
}
