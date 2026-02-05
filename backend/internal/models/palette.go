package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Palette struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID     uuid.UUID      `gorm:"type:uuid;index" json:"user_id"`
	ImageURL   string         `gorm:"type:text" json:"image_url"`
	Color1     string         `gorm:"type:varchar(7)" json:"color_1"`
	Color2     string         `gorm:"type:varchar(7)" json:"color_2"`
	Color3     string         `gorm:"type:varchar(7)" json:"color_3"`
	Color4     string         `gorm:"type:varchar(7)" json:"color_4"`
	Color5     string         `gorm:"type:varchar(7)" json:"color_5"`
	Name       string         `gorm:"type:varchar(100)" json:"name"`
	ShareCount int            `gorm:"default:0" json:"share_count"`
	CreatedAt  time.Time      `json:"created_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type PaletteStats struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"user_id"`
	TotalPalettes  int       `gorm:"default:0" json:"total_palettes"`
	TotalShares    int       `gorm:"default:0" json:"total_shares"`
	FavoriteColor  string    `gorm:"type:varchar(7)" json:"favorite_color"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
