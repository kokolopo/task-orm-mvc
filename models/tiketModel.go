package models

import "time"

type Tiket struct {
	NoTiket             int    `gorm:"type:int(25);not null"`
	AsalBerangkat       string `gorm:"type:varchar(100);not null"`
	Tujuan              string `gorm:"type:varchar(100);not null"`
	WaktuPemberangkatan time.Time
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
