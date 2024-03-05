package controllers

import (
	"miniProjectAPI/initializers"
	"miniProjectAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func QuizIndex(c *gin.Context) {
	var quiz []models.Quiz

	initializers.DB.Find(&quiz)
	c.JSON(http.StatusOK, gin.H{"quiz": quiz})
}

func QuizShow(c *gin.Context) {
	var Quiz models.Quiz
	id := c.Param("id")

	if err := initializers.DB.First(&Quiz, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"users": Quiz})
}

func QuizCreate(c *gin.Context) {

	var Quiz models.Quiz

	// Binding data JSON ke struct Quiz
	if err := c.ShouldBindJSON(&Quiz); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Membuat data Quiz baru
	initializers.DB.Create(&Quiz)
	c.JSON(http.StatusOK, gin.H{"message": "Quiz berhasil dibuat", "Quiz": Quiz})
}

func QuizUpdate(c *gin.Context) {

	var Quiz models.Quiz
	id := c.Param("id")

	// Binding data JSON ke struct User
	if err := c.ShouldBindJSON(&Quiz); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if initializers.DB.Model(&Quiz).Where("id = ?", id).Updates(&Quiz).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate Quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz berhasil diperbarui", "Quiz": Quiz})
}

func QuizDelete(c *gin.Context) {
	var quiz models.Quiz
	id := c.Param("id")

	// Cari quiz berdasarkan ID
	if err := initializers.DB.First(&quiz, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Quiz tidak ditemukan"})
		return
	}

	// Lakukan soft delete quiz dari database
	if err := initializers.DB.Delete(&quiz).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Gagal menghapus Quiz"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quiz berhasil dihapus"})
}
