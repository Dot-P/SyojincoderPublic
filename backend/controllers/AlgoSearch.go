package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "strconv"
)

func AlgoSearch(c *gin.Context) {
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

    fmt.Printf("Option as integer: %d\n", optionInt)

    algoResults, err := expAlgo(optionInt)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // 結果を直接JSON形式で返す
    c.JSON(http.StatusOK, algoResults)
}
