package main

import (
	"gogolook-interview/controllers"
	"gogolook-interview/repositorys"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	repositorys.DatabaseSetUp()

	r.GET("/tasks", controllers.GetTasks)
	r.POST("/task", controllers.CreateTask)
	r.PUT("/task/:id", controllers.UpdateTask)
	r.DELETE("/task/:id", controllers.DeleteTask)
	if err := r.Run("0.0.0.0:5566"); err != nil {
		panic(err)
	}
}
