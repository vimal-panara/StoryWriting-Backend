package controllers

import (
	"net/http"
	"story-plateform/config"
	"story-plateform/models"

	"github.com/gin-gonic/gin"
)

func GetStories(c *gin.Context) {
	var stories []models.Story

	result := config.DB.Find(&stories)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, stories)
}

func GetStoryById(c *gin.Context) {
	id := c.Param("id")
	var story models.Story

	result := config.DB.First(&story, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Story not found",
		})
		return
	}

	c.JSON(http.StatusOK, story)
}

func CreateStory(c *gin.Context) {
	var story models.Story

	if err := c.ShouldBindJSON(&story); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := config.DB.Create(&story)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, story)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, story)
}

func UpdateStory(c *gin.Context) {
	id := c.Param("id")
	var story models.Story

	if err := c.ShouldBindJSON(&story); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := config.DB.Model(&story).Where("id = ?", id).Updates(story)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, story)
}

func DeleteStory(c *gin.Context) {
	id := c.Param("id")
	var story models.Story

	result := config.DB.Delete(&story, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Story deleted successfully",
	})
}
