package controllers

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks := models.GetAllBooks()
	c.JSON(200, newBooks)
}

func GetBookById(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, _ := models.GetBookById(ID)
	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	CreateBook := &models.Book{}
	if err := c.BindJSON(&CreateBook); err != nil {
		log.Fatal("CreateBook:: Error in parsing", err)
	}
	book := CreateBook.CreateBook()

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(ID)
	c.JSON(200, book)
}

func UpdateBook(c *gin.Context) {
	var updateBook = &models.Book{}
	if err := c.BindJSON(&updateBook); err != nil {
		log.Fatal("UpdateBook:: Error in parsing", err)
	}
	bookId := c.Param("bookId")
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error is parsing id")
	}

	bookDetails, db := models.GetBookById(Id)
	if bookDetails.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if bookDetails.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)

	c.JSON(200, bookDetails)
}
