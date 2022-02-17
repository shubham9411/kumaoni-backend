package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

/*
Name always represent in english others have hindi, kfy prefix
*/
type Phrase struct {
	ID        string `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Content      string `gorm:"not null;size:1000" json:"content"`
	HindiContent string `gorm:"not null;size:1000" json:"hindiContent"`
	KfyContent   string `gorm:"not null;size:1000" json:"kfyContent"`
	AudioUrl     string `json:"audioUrl"`
	Approved     bool   `json:"approved" sql:"index"`

	CategoryId string   `gorm:"not null;default:'e2c450eb-b26c-49d4-8945-6d30e54dd2a6'" sql:"index" json:"categoryId"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
}

func (entity *Phrase) BeforeCreate(db *gorm.DB) error {
	entity.ID = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *Phrase) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}

func (w *Phrase) CreatePhrase() (*Phrase, error) {
	DB.NewRecord(w)
	if dbe := DB.Create(&w); dbe.Error != nil {
		return nil, dbe.Error
	}
	return w, nil
}

func GetAllPhrases() ([]Phrase, error) {
	var Phrases []Phrase
	dbe := DB.Where("APPROVED = true").Find(&Phrases)
	if dbe.Error != nil {
		return nil, dbe.Error
	}
	return Phrases, nil
}

func GetAllPhrasesByCategory(categoryId string) ([]Phrase, error) {
	var Phrases []Phrase
	dbe := DB.Where("APPROVED = true and CATEGORY_ID = ?", categoryId).Find(&Phrases)
	if dbe.Error != nil {
		return nil, dbe.Error
	}
	return Phrases, nil
}

func GetPhraseById(Id string) (*Phrase, *gorm.DB, error) {
	var getPhrase Phrase
	db := DB.Where("ID = ?", Id).Find(&getPhrase)
	if db.Error != nil {
		return nil, db, db.Error
	}
	return &getPhrase, db, nil
}

func UpdatePhrase(w *Phrase) (*Phrase, error) {
	db := DB.Model(Phrase{}).Where("id = ?", w.ID).Updates(w)
	if db.Error != nil {
		return nil, db.Error
	}
	return w, nil
}

func DeletePhrase(Id string) (*Phrase, error) {
	var phrase Phrase
	db := DB.Where("ID = ?", Id).Find(&phrase)
	if db.Error != nil {
		return nil, db.Error
	}

	db = db.Delete(phrase)
	if db.Error != nil {
		return nil, db.Error
	}
	return &phrase, nil
}
