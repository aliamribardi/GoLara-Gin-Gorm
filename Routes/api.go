package Routes

import (
	"GoLara-Gin-Gorm/Controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	Route := gin.Default()

	// User
	Route.GET("api/user", Controllers.IndexUser)
	Route.GET("api/user/:id", Controllers.ShowUser)
	Route.POST("api/user", Controllers.CreateUser)
	Route.PUT("api/user/:id", Controllers.UpdateUser)
	Route.DELETE("api/user/", Controllers.DeleteUser)

	// Kelas
	Route.GET("api/kelas", Controllers.IndexKelas)
	Route.GET("api/kelas/:id", Controllers.ShowKelas)
	Route.POST("api/kelas", Controllers.CreateKelas)
	Route.PUT("api/kelas/:id", Controllers.UpdateKelas)
	Route.DELETE("api/kelas", Controllers.DeleteKelas)

	Route.Run()
}