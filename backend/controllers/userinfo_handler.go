package controllers

import (
	"net/http"

	"github.com/dot_p/syojin/controllers/services"
	"github.com/dot_p/syojin/models"

	"github.com/gin-gonic/gin"
)

type UserInfoController struct {
	service services.UserInfoServicer
}

func NewUserInfoController(s services.UserInfoServicer) *UserInfoController {
	return &UserInfoController{service: s}
}

func (con *UserInfoController) GetUserInfoHandler(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	submissions, wrongs := con.service.GetSubmissionService(user.UserName)

	Performances, Rates, ContestNames := con.service.GetPerformanceService(user.UserName)

	Similarities, Recommends, ProcessedWrongs := con.service.CalEmbeddingService(wrongs)

	UserIdentifer := 2

	if len(Rates) == 0 || Rates[len(Rates)-1] == 0 {
		UserIdentifer = 1
	}

	data := models.AtcoderUserInfo{
		UserName:        user.UserName,
		UserIdentifer:   UserIdentifer,
		Submissions:     submissions,
		Performances:    Performances,
		Rates:           Rates,
		ContestNames:    ContestNames,
		Similarities:    Similarities,
		Recommends:      Recommends,
		ProcessedWrongs: ProcessedWrongs,
	}

	c.JSON(http.StatusOK, data)
}
