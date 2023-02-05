package main

import (
	"gogolook-interview/repositorys"
	"gogolook-interview/router"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	repositorys.DatabaseSetUp()

	r := router.Api()
	r.Run(":5566")
}
