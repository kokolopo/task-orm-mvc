package models

import "time"

type Transaksi struct {
	Id        int   `gorm:"primary_key;auto_increment;not_null"`
	TiketId   int   `gorm:"type:int(25);not null"`
	Tiket     Tiket `gorm:"foreignKey:TiketId;not null"`
	UserId    int   `gorm:"type:int(25);not null"`
	User      User  `gorm:"foreignKey:UserId;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
