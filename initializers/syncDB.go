package initializers

import (
	"miniProjectAPI/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Quiz{}, &models.Pertanyaan{}, &models.JawabanPeserta{})
}
