package Controllers

import (
	"encoding/json"
	"net/http"
	"GoLara-Gin-Gorm/Models"
	"GoLara-Gin-Gorm/Database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func IndexUser(c *gin.Context) {
	var users []Models.User

	Database.DB.Preload("Kelas").Find(&users)
	if users == nil {
		c.JSON(http.StatusBadRequest, gin.H{"users": "Belum Ada Data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var user Models.User
	
	err := c.ShouldBindJSON(&user);
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	Database.DB.Preload("Kelas").Create(&user)
	
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func ShowUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")

	if err := Database.DB.Preload("Kelas").First(&user, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if Database.DB.Model(&user).Where("id = ?", id).Updates(&user).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat mengupdate User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Diupdate"})
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	
	var input struct {
		Id	json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	id, _ := input.Id.Int64()
	if Database.DB.Delete(&user, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus User"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})
}