package models

import "time"

type User struct {
	Id        int    `gorm:"primary_key;auto_increment;not_null"`
	Nama      string `gorm:"type:varchar(100);not null"`
	Usia      int    `gorm:"type:int(25);not null"`
	NoTelpn   string `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
