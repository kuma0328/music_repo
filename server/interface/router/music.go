package router

import (
	"github.com/kuma0328/music_repo/infrastructure/database"
	"github.com/kuma0328/music_repo/infrastructure/persistance"
	"github.com/kuma0328/music_repo/interface/handler"
	"github.com/kuma0328/music_repo/usecase"
)


func (r Router) InitMusicRouter(conn *database.Conn) {
	repo := persistance.NewMusicRepositry(conn)
	uc := usecase.NewMusicUsecase(repo)
	h := handler.NewMusicHandler(uc)

	
}