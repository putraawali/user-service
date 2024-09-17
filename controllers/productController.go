package controllers

import (
	"net/http"
	"user-service/database"
	"user-service/helpers"
	"user-service/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	product := models.Product{}
	userId := uint(userData["id"].(float64))

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&product)
	} else {
		err = c.ShouldBind(&product)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	product.UserID = userId

	if err = db.Debug().Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, product)
}
