package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ningenme/nina-api/pkg/infra"
	"net/http"
)

type Controller struct{}

func (Controller) GetHealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
	})
}
func (Controller) GetGithubContributionList(c *gin.Context) {
	repository := infra.ContributionRepository{}
	c.JSON(http.StatusOK, gin.H{
		"contributionList": repository.GetList(),
	})
}
