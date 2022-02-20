package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
	"github.com/shubham9411/kumaoni-backend/utils"
)

var NewPhrase models.Phrase

func GetAllPhrases(c *gin.Context) {
	newPhrases, err := models.GetAllPhrases()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, newPhrases)
}

func GetAllPhrasesByCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")
	newPhrases, err := models.GetAllPhrasesByCategory(categoryId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, newPhrases)
}

func GetPhraseById(c *gin.Context) {
	phraseId := c.Param("phraseId")

	phrase, _, err := models.GetPhraseById(phraseId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, phrase)
}

func CreatePhrase(c *gin.Context) {
	CreatePhrase := &models.Phrase{}
	if err := c.BindJSON(&CreatePhrase); err != nil {
		utils.SendError("Error in parsing", c)
		return
	}
	phrase, err := CreatePhrase.CreatePhrase()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, phrase)
}

func DeletePhrase(c *gin.Context) {
	phraseId := c.Param("phraseId")
	phrase, err := models.DeletePhrase(phraseId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, phrase)
}

func UpdatePhrase(c *gin.Context) {
	var updatePhrase = &models.Phrase{}
	if err := c.BindJSON(&updatePhrase); err != nil {
		utils.SendError("Error in parsing", c)
		return
	}
	updatePhrase.ID = c.Param("phraseId")
	phraseDetails, err := models.UpdatePhrase(updatePhrase)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, phraseDetails)
}
