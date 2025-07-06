package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Anime struct {
	ID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title  string
	Studio string
	Years  string
	Season string
	gorm.Model
}
