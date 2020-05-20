package routes

import (
	"../controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.IndexHome)
	router.GET("/index", controllers.IndexHome)

	router.GET("/sahibinden", controllers.Sahibinden)
	router.GET("/fiyatlar", controllers.GetPriceById)

}
