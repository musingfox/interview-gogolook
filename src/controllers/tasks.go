package controllers

import (
	"net/http"
	"task/requests"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

func GetTasks(c *gin.Context) {
	var mockTasks []Task

	/* get task from DB */

	mockTasks = append(mockTasks, Task{
		ID:     1,
		Name:   "name",
		Status: false,
	})

	c.JSON(http.StatusOK, gin.H{"result": mockTasks})
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
	mockTask := Task{
		ID:     1,
		Name:   req.Name,
		Status: false,
	}

	c.JSON(http.StatusOK, mockTask)
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

	/* update task to DB */
	mockTask := Task{
		ID:     1,
		Name:   req.Name,
		Status: req.Status,
	}

	c.JSON(http.StatusCreated, mockTask)
}
func DeleteTask(c *gin.Context) {
	var req requests.DeleteTaskRequest

	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	/* delete task to DB */

	c.JSON(http.StatusOK, nil)
}
