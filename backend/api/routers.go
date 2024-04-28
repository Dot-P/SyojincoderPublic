package api

import (
	"database/sql"

	"github.com/dot_p/syojin/controllers"
	"github.com/dot_p/syojin/services"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	ser := services.NewMyAppService(db)
	userCon := controllers.NewUserInfoController(ser)
	algoCon := controllers.NewAlgoSearchController(ser)

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/hello", controllers.HealthCheck)
	router.POST("/dashboard", userCon.GetUserInfoHandler)
	router.POST("/algo_search", algoCon.AlgoSearchHandler)

	return router
}
