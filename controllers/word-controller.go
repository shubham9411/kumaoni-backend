package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
)

var NewWord models.Word

func GetAllWords(c *gin.Context) {
	newWords := models.GetAllWords()
	c.JSON(200, newWords)
}

func GetWordById(c *gin.Context) {
	wordId := c.Param("wordId")
	ID, err := strconv.ParseInt(wordId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	word, _ := models.GetWordById(ID)
	c.JSON(200, word)
}

func CreateWord(c *gin.Context) {
	CreateWord := &models.Word{}
	if err := c.BindJSON(&CreateWord); err != nil {
		log.Fatal("CreateWord:: Error in parsing", err)
	}
	word := CreateWord.CreateWord()

	c.JSON(200, word)
}

func DeleteWord(c *gin.Context) {
	wordId := c.Param("wordId")
	ID, err := strconv.ParseInt(wordId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	word := models.DeleteWord(ID)
	c.JSON(200, word)
}

func UpdateWord(c *gin.Context) {
	var updateWord = &models.Word{}
	if err := c.BindJSON(&updateWord); err != nil {
		log.Fatal("UpdateWord:: Error in parsing", err)
	}
	wordDetails := models.UpdateWord(updateWord)

	c.JSON(200, wordDetails)
}
