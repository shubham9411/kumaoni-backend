package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

/*
Name always represent in english others have hindi, kfy prefix
*/
type Word struct {
	gorm.Model
	Content      string `gorm:"unique;not null;default:'A'" json:"content"`
	HindiContent string `gorm:"not null;" json:"hindiContent"`
	KfyContent   string `gorm:"not null;" json:"kfyContent"`
	Origin       string `json:"origin"`
	AudioUrl     string `json:"audioUrl"`
	Phonetic     string `json:"phonetic"`
	Type         string `json:"type"`
	Usage        string `gorm:"type:text" json:"usage"`
	ImageUrl     string `json:"imageUrl"`

	CategoryId uint     `gorm:"not null" sql:"index" json:"categoryId"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET 1;" json:"-"`
}

func (b *Word) CreateWord() *Word {
	DB.NewRecord(b)
	DB.Create(&b)
	return b
}

func GetAllWords() []Word {
	var Words []Word
	DB.Find(&Words)
	return Words
}

func GetWordById(Id int64) (*Word, *gorm.DB) {
	var getWord Word
	db := DB.Where("ID = ?", Id).Find(&getWord)
	return &getWord, db
}

func UpdateWord(w *Word) *Word {
	err := DB.Model(Word{}).Where("id = ?", w.ID).Updates(w)
	if err != nil {
		fmt.Println(err)
	}
	return w
}

func DeleteWord(Id int64) Word {
	var word Word
	DB.Where("ID = ?", Id).Delete(word)
	return word
}
