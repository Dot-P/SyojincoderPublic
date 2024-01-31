package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "Syojincoder2/models"
)

func GetUserName(c *gin.Context) {
    var user models.User

    if err := c.BindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    submissions, wrongs := getSubmission(user.UserName)
    Performances, Rates, ContestNames := getPerformance(user.UserName)
    Similarities, Recommends, ProcessedWrongs := calEmbedding(wrongs)

    UserIdentifer := 2

    if len(Rates) == 0 || Rates[len(Rates)-1] == 0{
        UserIdentifer = 1
    }

    type ResponseData struct {
        UserName      string   `json:"userName"`
        UserIdentifer int  `json:"UserIdentifier"`
        Submissions   []int    `json:"submissions"`
        Performances  []int    `json:"performances"`
        Rates         []int    `json:"rates"`
        ContestNames  []string `json:"contestNames"`
        Similarities  []float64 `json:"similarities"`
        Recommends    []string `json:"recommends"`
        ProcessedWrongs    []string `json:"processedWrongs"`
    }

    data := ResponseData{
        UserName:      user.UserName,
        UserIdentifer: UserIdentifer,
        Submissions:   submissions,
        Performances:  Performances,
        Rates:         Rates,
        ContestNames:  ContestNames,
        Similarities:  Similarities,
        Recommends:    Recommends,
        ProcessedWrongs: ProcessedWrongs,
    }

    c.JSON(http.StatusOK, data)
}
