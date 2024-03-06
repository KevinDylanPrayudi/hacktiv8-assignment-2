package controller

import (
	"assigment-2/database"
	"assigment-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	appJSON = "application/json"
)

func GetOrders(ctx *gin.Context) {
	db := database.GetDB()
	var orders = []models.Order{}
	db.Model(&models.Order{}).Preload("Items").Find(&orders)
	ctx.JSON(http.StatusOK, orders)
}

func CreateOrder(ctx *gin.Context) {
	order := models.Order{}
	ctx.ShouldBind(&order)
	db := database.GetDB()
	db.Create(&order)
	ctx.JSON(http.StatusCreated, order)
}

func UpdateOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "The parameter is invalid",
		})
		return
	}
	order := models.Order{}
	db := database.GetDB()
	ctx.ShouldBind(&order)
	order.ID = uint(id)
	err = db.Where("id = ?", id).Session(&gorm.Session{FullSaveAssociations: true}).Updates(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "The order isn't found",
		})
		return
	}

	ctx.JSON(http.StatusOK, order)
}

func DeleteOrder(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "The parameter is invalid",
		})
		return
	}
	order := models.Order{}
	db := database.GetDB()
	order.ID = uint(id)
	err = db.Where("id = ?", id).Delete(&order).Error
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Error occured in the server",
		})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
