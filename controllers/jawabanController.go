package controllers

import (
	"miniProjectAPI/initializers"
	"miniProjectAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func JawabanIndex(c *gin.Context) {

	var jawaban_peserta []models.JawabanPeserta

	initializers.DB.Find(&jawaban_peserta)
	c.JSON(http.StatusOK, gin.H{"jawaban_peserta": jawaban_peserta})
}

func JawabanShow(c *gin.Context) {

	var jawaban_peserta models.JawabanPeserta
	id := c.Param("id")

	if err := initializers.DB.First(&jawaban_peserta, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": jawaban_peserta})
}

func JawabanCreate(c *gin.Context) {

	var jawaban_peserta models.JawabanPeserta

	// Binding data JSON ke struct jawaban_peserta
	if err := c.ShouldBindJSON(&jawaban_peserta); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Membuat data jawaban_peserta baru
	initializers.DB.Create(&jawaban_peserta)
	c.JSON(http.StatusOK, gin.H{"message": "Jawaban berhasil di input", "jawaban": jawaban_peserta})
}

func JawabanUpdate(c *gin.Context) {

	var jawaban_peserta models.JawabanPeserta
	id := c.Param("id")

	// Binding data JSON ke struct User
	if err := c.ShouldBindJSON(&jawaban_peserta); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Model(&jawaban_peserta).Where("id = ?", id).Updates(&jawaban_peserta).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate jawaban peserta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "jawaban peserta berhasil diperbarui", "jawaban_peserta": jawaban_peserta})
}

func JawabanDelete(c *gin.Context) {
	var jawaban_peserta models.JawabanPeserta
	id := c.Param("id")

	// Cari user berdasarkan ID
	if err := initializers.DB.First(&jawaban_peserta, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data jawaban tidak ditemukan"})
		return
	}

	// Hapus user dari database
	if err := initializers.DB.Delete(&jawaban_peserta).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus jawaban peserta"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "jawaban peserta berhasil dihapus"})
}
