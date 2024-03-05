package models

import (
	"time"

	"gorm.io/gorm"
)

type Quiz struct {
	gorm.Model
	Judul        string    `gorm:"type:varchar(255);not null" json:"judul"`
	Deskripsi    string    `gorm:"type:text" json:"deskripsi"`
	WaktuMulai   time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"waktu_mulai"`
	WaktuSelesai time.Time `gorm:"type:timestamp;not null;default:current_timestamp" json:"waktu_selesai"`
}
