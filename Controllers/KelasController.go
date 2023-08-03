package Controllers

import (
	"encoding/json"
	"net/http"
	"GoLara-Gin-Gorm/Models"
	"GoLara-Gin-Gorm/Database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexKelas(c *gin.Context) {
	var kelas []Models.Kelas

	Database.DB.Find(&kelas)
	if kelas == nil {
		c.JSON(http.StatusBadRequest, gin.H{"kelas": "Belum Ada Data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"kelas": kelas})
}

func CreateKelas(c *gin.Context) {
	var kelas Models.Kelas
	
	err := c.ShouldBindJSON(&kelas);
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	Database.DB.Create(&kelas)
	c.JSON(http.StatusOK, gin.H{"kelas": kelas})
}

func ShowKelas(c *gin.Context) {
	var kelas Models.Kelas
	id := c.Param("id")

	if err := Database.DB.First(&kelas, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"kelas": kelas})
}

func UpdateKelas(c *gin.Context) {
	var kelas Models.Kelas
	id := c.Param("id")

	if err := c.ShouldBindJSON(&kelas); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if Database.DB.Model(&kelas).Where("id = ?", id).Updates(&kelas).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diupdate"})
}

func DeleteKelas(c *gin.Context) {
	var kelas Models.User
	
	var input struct {
		Id	json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	id, _ := input.Id.Int64()
	if Database.DB.Delete(&kelas, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})
}