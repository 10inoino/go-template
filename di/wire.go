//go:build wireinject
// +build wireinject

package di

import (
	"example/web-service-gin/db"
	"example/web-service-gin/src/presentation/controller"
	"example/web-service-gin/src/repository/postgres/repository"

	album_uc "example/web-service-gin/src/usecase/album"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitializeEvent() (*gin.Engine, error) {
	wire.Build(
		controller.NewAlbumController,
		controller.NewHealthCheckController,
		repository.NewAlbumRepository,
		album_uc.NewCreateAlbumUsecase,
		album_uc.NewGetAlbumUsecase,
		album_uc.NewListAlbumUsecase,
		album_uc.NewUpdateAlbumUsecase,
		album_uc.NewDeleteAlbumUsecase,
		db.NewPostgresDB,
		NewGinEngine,
	)
	return nil, nil
}
