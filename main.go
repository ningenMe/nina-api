package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ningenme/nina-api/pkg/controller"
)

func main() {
	r := gin.Default()
	co := controller.Controller{}
	r.GET("/healthcheck", co.GetHealthCheck)
	r.GET("/github-contributions", co.GetGithubContributionList)
	r.Run(":8081")
}
