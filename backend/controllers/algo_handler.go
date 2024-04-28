package controllers

import (
	"net/http"
	"strconv"

	"github.com/dot_p/syojin/controllers/services"
	"github.com/gin-gonic/gin"
)

type AlgoSearchController struct {
	service services.AlgoSearchServicer
}

func NewAlgoSearchController(s services.AlgoSearchServicer) *AlgoSearchController {
	return &AlgoSearchController{service: s}
}

func (con *AlgoSearchController) AlgoSearchHandler(c *gin.Context) {
	var requestData map[string]interface{}

	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	optionValue, ok := requestData["option"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Option not found"})
		return
	}

	var optionInt int
	switch v := optionValue.(type) {
	case float64:
		optionInt = int(v)
	case int:
		optionInt = v
	case string:
		var err error
		optionInt, err = strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Option is not a valid number"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Option is of an invalid type"})
		return
	}

	algoResults, err := con.service.AlgoSearchServicer(optionInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 結果を直接JSON形式で返す
	c.JSON(http.StatusOK, algoResults)
}
