package models

import "gorm.io/gorm"

type Pertanyaan struct {
	gorm.Model
	Pertanyaan   string `gorm:"type:text;not null" json:"pertanyaan"`
	OpsiJawaban  string `gorm:"type:text;not null" json:"opsi_jawaban"`
	JawabanBenar int    `gorm:"not null" json:"jawaban_benar"`
	IdQuiz       uint   `gorm:"not null;foreignKey:IdQuiz" json:"id_quiz"`
}
