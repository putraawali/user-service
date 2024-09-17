package controllers

import (
	"errors"
	"net/http"
	"user-service/database"
	"user-service/helpers"
	"user-service/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var appJSON = "application/json"

func UserRegister(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	user := models.User{}

	var err error

	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := db.Debug().Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal server error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":        user.ID,
		"email":     user.Email,
		"full_name": user.FullName,
	})
}

func UserLogin(c *gin.Context) {
	db := database.GetDB()

	contentType := helpers.GetContentType(c)

	user := models.User{}

	password := ""

	var err error
	if contentType == appJSON {
		err = c.ShouldBindBodyWithJSON(&user)
	} else {
		err = c.ShouldBind(&user)
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	password = user.Password

	if err = db.Debug().First(&user, "email = ?", user.Email).Error; err != nil {
		statusCode := http.StatusInternalServerError
		msg := err.Error()

		if errors.Is(err, gorm.ErrRecordNotFound) {
			statusCode = http.StatusNotFound
			msg = "Email not found"
		}

		c.JSON(statusCode, gin.H{
			"error":   http.StatusText(statusCode),
			"message": msg,
		})
		return
	}

	if !helpers.ComparePassword([]byte(user.Password), []byte(password)) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   http.StatusText(http.StatusBadRequest),
			"message": "Wrong password",
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
