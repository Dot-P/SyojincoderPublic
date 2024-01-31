package router

import(
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"Syojincoder2/controllers"
)

func SetupRouter() *gin.Engine{

	router := gin.Default()

	router.Use(cors.Default())

	router.GET("/test", controllers.Test)
	router.POST("/user_name", controllers.GetUserName)
	router.POST("/algo_search", controllers.AlgoSearch)

	return router
}