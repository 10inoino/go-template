package main

import (
	"database/sql"
	"example/web-service-gin/src/presentation/controller"
	"example/web-service-gin/src/repository/postgres/repository"
	"fmt"
	"os"

	album_uc "example/web-service-gin/src/usecase/album"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// TODO:他のディレクトリに移動
func setUpPostgres() (*sql.DB, error) {
	host := os.Getenv("PSQL_HOST")
	dbname := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", host, dbname, user, password))
	return db, err
}

func setUpMySQL() (*sql.DB, error) {
	dsn := os.Getenv("DSN")
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s&charset=utf8&parseTime=true", dsn))
	return db, err
}

func main() {
	app := gin.Default()
	db, dbErr := setUpMySQL()
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

	app.GET("/albums", albumCon.ListAlbums)
	app.GET("/albums/:id", albumCon.GetAlbumByID)
	app.POST("/albums", albumCon.CreateAlbum)
	app.PUT("/albums", albumCon.UpdateAlbum)
	app.DELETE("/albums/:id", albumCon.DeleteAlbum)
	app.GET("/health", healthCheckCon.HealthCheck)

	app.Run(":8080")
}
