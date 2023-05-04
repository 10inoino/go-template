package main

import (
	"database/sql"
	"example/web-service-gin/src/presentation/controller"
	"example/web-service-gin/src/repository/postgres/repository"
	"example/web-service-gin/src/usecase"
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
	createAlbumUsecase := usecase.NewCreateAlbumUsecase(albumRepo)
	getAlbumUsecase := usecase.NewGetAlbumUsecase(albumRepo)
	listAlbumUsecase := usecase.NewListAlbumUsecase(albumRepo)
	updateAlbumUsecase := usecase.NewUpdateAlbumUsecase(albumRepo)
	deleteAlbumUsecase := usecase.NewDeleteAlbumUsecase(albumRepo)
	albumCon := controller.NewAlbumController(
		*createAlbumUsecase,
		*getAlbumUsecase,
		*listAlbumUsecase,
		*updateAlbumUsecase,
		*deleteAlbumUsecase,
	)

	router.GET("/albums", albumCon.ListAlbums)
	router.GET("/albums/:id", albumCon.GetAlbumByID)
	router.POST("/albums", albumCon.CreateAlbum)
	router.PUT("/albums", albumCon.UpdateAlbum)
	router.DELETE("/albums/:id", albumCon.DeleteAlbum)

	router.Run("localhost:8080")
}
