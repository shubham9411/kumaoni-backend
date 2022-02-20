package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/shubham9411/kumaoni-backend/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string     `validate:"required" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`

	FirstName      string `validate:"required" gorm:"not null" json:"first_name"`
	LastName       string `validate:"required" gorm:"not null" json:"last_name"`
	HashedPassword string `gorm:"not null" json:"-"`
	Password       string `validate:"required,gte=8" gorm:"-" json:"password"`
	Email          string `validate:"required,email" gorm:"unique;not null" json:"email"`
	Phone          string `validate:"required,gte=10,lte=12" json:"phone"`
	Token          string `validate:"required" gorm:"size:1000" json:"token"`
	RefreshToken   string `validate:"required" gorm:"size:1000" json:"refresh_token"`
	UserType       string `validate:"required,oneof=USER ADMIN" json:"user_type"`
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func HashPassword(password string) string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(pass)
}

func (user *User) CreateUser() (*User, error) {
	DB.NewRecord(user)
	user.ID = uuid.New().String()
	user.UserType = "USER"
	token, refreshToken, _ := utils.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.UserType, user.ID)
	user.Token = token
	user.RefreshToken = refreshToken

	err := validate.Struct(user)

	if err != nil {
		return nil, err
	}

	password := HashPassword(user.Password)
	user.HashedPassword = password
	if dbe := DB.Create(&user); dbe.Error != nil {
		return nil, dbe.Error
	}
	user.Password = ""
	return user, nil
}

func UpdateTokens(userId string, token string, refreshToken string) {
	user, err := GetUserById(userId)
	if err != nil {
		log.Panic(err)
	}
	user.Token = token
	user.RefreshToken = refreshToken
	DB.Model(&user).Updates(User{Token: token, RefreshToken: refreshToken})
}

func GetAllUsers() ([]User, error) {
	var users []User
	db := DB.Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	return users, nil
}

func GetUserById(userId string) (*User, error) {
	var user User
	db := DB.Where("ID = ?", userId).Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	db := DB.Where("email = ? ", email).Find(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}
