package models

import "time"

type ShortLink struct {
	ID           uint
	OriginalURL  string `gorm:"type:longtext;not null"`
	ShortenedURL string `gorm:"type:varchar(100);unique;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
