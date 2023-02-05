package main

import (
	"task/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/task", controllers.CreateTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)
	if err := r.Run("localhost:5566"); err != nil {
		panic(err)
	}
}
