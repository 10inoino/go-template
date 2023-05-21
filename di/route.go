package di

import (
	"example/web-service-gin/src/presentation/controller"

	"github.com/gin-gonic/gin"
)

func NewGinEngine(
	albumCon controller.AlbumController,
	healthCheckCon controller.HealthCheckController,
) *gin.Engine {
	app := gin.New()

	app.GET("/albums", albumCon.ListAlbums)
	app.GET("/albums/:id", albumCon.GetAlbumByID)
	app.POST("/albums", albumCon.CreateAlbum)
	app.PUT("/albums", albumCon.UpdateAlbum)
	app.DELETE("/albums/:id", albumCon.DeleteAlbum)
	app.GET("/health", healthCheckCon.HealthCheck)

	return app
}
