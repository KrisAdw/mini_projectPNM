package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama     string `gorm:"type:varchar(255);not null" json:"nama_user"`
	Email    string `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Role     string `gorm:"type:varchar(50);not null" json:"role"`
}
