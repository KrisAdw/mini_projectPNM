package controllers

import (
	"miniProjectAPI/initializers"
	"miniProjectAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PertanyaanIndex(c *gin.Context) {
	var Pertanyaan []models.Pertanyaan

	initializers.DB.Find(&Pertanyaan)
	c.JSON(http.StatusOK, gin.H{"pertanyaan": Pertanyaan})
}

func PertanyaanShow(c *gin.Context) {
	var Pertanyaan models.Pertanyaan
	id := c.Param("id")

	if err := initializers.DB.First(&Pertanyaan, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": Pertanyaan})
}

func PertanyaanCreate(c *gin.Context) {
	var Pertanyaan models.Pertanyaan

	// Binding data JSON ke struct Pertanyaan
	if err := c.ShouldBindJSON(&Pertanyaan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Membuat data Pertanyaan baru
	initializers.DB.Create(&Pertanyaan)
	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan berhasil dibuat", "Pertanyaan": Pertanyaan})
}

func PertanyaanUpdate(c *gin.Context) {
	var Pertanyaan models.Pertanyaan
	id := c.Param("id")

	// Binding data JSON ke struct User
	if err := c.ShouldBindJSON(&Pertanyaan); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Model(&Pertanyaan).Where("id = ?", id).Updates(&Pertanyaan).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate Pertanyaan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan berhasil diperbarui", "Pertanyaan": Pertanyaan})
}

func PertanyaanDelete(c *gin.Context) {
	var Pertanyaan models.Pertanyaan
	id := c.Param("id")

	// Cari user berdasarkan ID
	if err := initializers.DB.First(&Pertanyaan, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Pertanyaan tidak ditemukan"})
		return
	}

	// Hapus user dari database
	if err := initializers.DB.Delete(&Pertanyaan).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus Pertanyaan"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Pertanyaan berhasil dihapus"})
}
