package models

import "gorm.io/gorm"

type JawabanPeserta struct {
	gorm.Model
	IdUser         uint `gorm:"foreignKey:IdUser;not null" json:"id_user"`
	IdQuiz         uint `gorm:"foreignKey:IdQuiz;not null" json:"id_quiz"`
	IdPertanyaan   uint `gorm:"foreignKey:IdPertanyaan;not null" json:"id_pertanyaan"`
	JawabanPeserta uint `gorm:"not null" json:"jawaban_peserta"`
	Skor           uint `gorm:"not null" json:"skor_peserta"`
}
