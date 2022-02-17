package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

/*
Name always represent in english others have hindi, kfy prefix
*/
type Word struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Content      string `gorm:"unique;not null;default:'A'" json:"content"`
	HindiContent string `gorm:"not null;" json:"hindiContent"`
	KfyContent   string `gorm:"not null;" json:"kfyContent"`
	Origin       string `json:"origin"`
	AudioUrl     string `json:"audioUrl"`
	Phonetic     string `json:"phonetic"`
	Type         string `json:"type"`
	Usage        string `gorm:"type:text" json:"usage"`
	ImageUrl     string `json:"imageUrl"`
	Approved     bool   `json:"approved" sql:"index"`

	CategoryId string   `gorm:"not null;default:'e2c450eb-b26c-49d4-8945-6d30e54dd2a6'" sql:"index" json:"categoryId"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (entity *Word) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Word) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (w *Word) CreateWord() (*Word, error) {
	DB.NewRecord(w)
	if dbe := DB.Create(&w); dbe.Error != nil {
		return nil, dbe.Error
	}
	return w, nil
}

func GetAllWords() ([]Word, error) {
	var Words []Word
	dbe := DB.Where("APPROVED = true").Find(&Words)
	if dbe.Error != nil {
		return nil, dbe.Error
	}
	return Words, nil
}

func GetAllWordsByCategory(categoryId string) ([]Word, error) {
	var Words []Word
	dbe := DB.Where("APPROVED = true and CATEGORY_ID = ?", categoryId).Find(&Words)
	if dbe.Error != nil {
		return nil, dbe.Error
	}
	return Words, nil
}

func GetWordById(Id string) (*Word, *gorm.DB, error) {
	var getWord Word
	db := DB.Where("ID = ?", Id).Find(&getWord)
	if db.Error != nil {
		return nil, db, db.Error
	}
	return &getWord, db, nil
}

func UpdateWord(w *Word) (*Word, error) {
	db := DB.Model(Word{}).Where("id = ?", w.ID).Updates(w)
	if db.Error != nil {
		return nil, db.Error
	}
	return w, nil
}

func DeleteWord(Id string) (*Word, error) {
	var word Word
	db := DB.Where("ID = ?", Id).Find(&word)
	if db.Error != nil {
		return nil, db.Error
	}

	db = db.Delete(word)
	if db.Error != nil {
		return nil, db.Error
	}
	return &word, nil
}
