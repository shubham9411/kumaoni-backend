package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/shubham9411/kumaoni-backend/utils"
)

type User struct {
	ID        string     `gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`

	FirstName    string `gorm:"not null" json:"first_name"`
	LastName     string `gorm:"not null" json:"last_name"`
	Password     string `gorm:"not null" json:"password"`
	Email        string `gorm:"unique;not null" json:"email"`
	Phone        string `json:"phone"`
	Token        string `gorm:"size:1000" json:"token"`
	RefreshToken string `gorm:"size:1000" json:"refresh_token"`
	UserType     string `json:"user_type" validate:"required, eq=ADMIN|eq=USER"`
}

func (entity *User) BeforeCreate(db *gorm.DB) error {
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *User) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (user *User) CreateUser() (*User, error) {
	DB.NewRecord(user)
	user.ID = uuid.New().String()
	token, refreshToken, _ := utils.GenerateAllTokens(user.Email, user.FirstName, user.LastName, user.UserType, user.ID)
	user.Token = token
	user.RefreshToken = refreshToken

	if dbe := DB.Create(&user); dbe.Error != nil {
		return nil, dbe.Error
	}
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
