package router

import (
	"gogolook-interview/controllers"

	"github.com/gin-gonic/gin"
)

func Api() *gin.Engine {
	r := gin.Default()

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/task", controllers.CreateTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)

	return r
}
