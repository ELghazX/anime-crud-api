package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Anime struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title  string
	Studio string
	Years  string
	Season string
	gorm.Model
}
