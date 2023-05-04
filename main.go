package main

import (
	"database/sql"
	"example/web-service-gin/src/presentation/controller"
	"example/web-service-gin/src/repository/postgres/repository"
	album_uc "example/web-service-gin/src/usecase/album"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func generateDB() (*sql.DB, error) {
	host := os.Getenv("PSQL_HOST")
	dbname := os.Getenv("PSQL_DBNAME")
	user := os.Getenv("PSQL_USER")
	password := os.Getenv("PSQL_PASS")

	return sql.Open(
		"postgres",
		fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", host, dbname, user, password))
}

func main() {
	router := gin.Default()
	db, dbErr := generateDB()
	if dbErr != nil {
		panic("failed database connection")
	}
	albumRepo := repository.NewAlbumRepository(db)
	createAlbumUsecase := album_uc.NewCreateAlbumUsecase(albumRepo)
	getAlbumUsecase := album_uc.NewGetAlbumUsecase(albumRepo)
	listAlbumUsecase := album_uc.NewListAlbumUsecase(albumRepo)
	updateAlbumUsecase := album_uc.NewUpdateAlbumUsecase(albumRepo)
	deleteAlbumUsecase := album_uc.NewDeleteAlbumUsecase(albumRepo)
	albumCon := controller.NewAlbumController(
		*createAlbumUsecase,
		*getAlbumUsecase,
		*listAlbumUsecase,
		*updateAlbumUsecase,
		*deleteAlbumUsecase,
	)
	healthCheckCon := controller.NewHealthCheckController()

	router.GET("/albums", albumCon.ListAlbums)
	router.GET("/albums/:id", albumCon.GetAlbumByID)
	router.POST("/albums", albumCon.CreateAlbum)
	router.PUT("/albums", albumCon.UpdateAlbum)
	router.DELETE("/albums/:id", albumCon.DeleteAlbum)
	router.GET("/health", healthCheckCon.HealthCheck)

	router.Run("localhost:8080")
}
