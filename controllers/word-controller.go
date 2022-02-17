package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
	"github.com/shubham9411/kumaoni-backend/utils"
)

var NewWord models.Word

func GetAllWords(c *gin.Context) {
	newWords, err := models.GetAllWords()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(200, newWords)
}

func GetAllWordsByCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")
	newWords, err := models.GetAllWordsByCategory(categoryId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(200, newWords)
}

func GetWordById(c *gin.Context) {
	wordId := c.Param("wordId")

	word, _, err := models.GetWordById(wordId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(200, word)
}

func CreateWord(c *gin.Context) {
	CreateWord := &models.Word{}
	if err := c.BindJSON(&CreateWord); err != nil {
		utils.SendError("Error in parsing", c)
		return
	}
	word, err := CreateWord.CreateWord()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(200, word)
}

func DeleteWord(c *gin.Context) {
	wordId := c.Param("wordId")
	word, err := models.DeleteWord(wordId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(200, word)
}

func UpdateWord(c *gin.Context) {
	var updateWord = &models.Word{}
	if err := c.BindJSON(&updateWord); err != nil {
		utils.SendError("Error in parsing", c)
		return
	}
	updateWord.ID = c.Param("wordId")
	wordDetails, err := models.UpdateWord(updateWord)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(200, wordDetails)
}
