package controllers

import (
	"net/http"
	"story-plateform/config"
	"story-plateform/models"
	"story-plateform/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user models.User

	if err := config.DB.Model(&user).Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if !utils.CheckHashedPassword(input.Password, user.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJwtToken(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", c.Request.RemoteAddr, false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", c.Request.RemoteAddr, false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully logged out",
	})
}
