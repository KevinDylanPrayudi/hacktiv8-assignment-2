package router

import (
	"assigment-2/controller"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	router := gin.Default()
	router.GET("/orders", controller.GetOrders)
	router.POST("/orders", controller.CreateOrder)
	router.PUT("/orders/:id", controller.UpdateOrder)
	router.DELETE("/orders/:id", controller.DeleteOrder)

	return router
}
