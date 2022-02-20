package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shubham9411/kumaoni-backend/models"
	"github.com/shubham9411/kumaoni-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

var Users models.User

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = "Email or Password is incorrect"
		check = false
	}
	return check, msg
}

func Signup(c *gin.Context) {
	user := &models.User{}
	if err := c.BindJSON(&user); err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	user, err := user.CreateUser()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundUser, err := models.GetUserByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isValidPassword, msg := VerifyPassword(user.Password, foundUser.HashedPassword)

	if !isValidPassword {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	token, refreshToken, _ := utils.GenerateAllTokens(foundUser.Email, foundUser.FirstName, foundUser.LastName, foundUser.UserType, foundUser.ID)
	models.UpdateTokens(foundUser.ID, token, refreshToken)

	c.JSON(http.StatusOK, foundUser)

}

func GetAllUsers(c *gin.Context) {
	userId := c.Param("user_id")

	if err := utils.CheckUserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := utils.MatchUserTypeToUid(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users, err := models.GetAllUsers()
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("user_id")
	if err := utils.MatchUserTypeToUid(c, userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserById(userId)
	if err != nil {
		utils.SendError(err.Error(), c)
		return
	}
	c.JSON(http.StatusOK, user)
}
