package usecase

import (
	"fmt"

	"github.com/kuma0328/music_repo/domain/entity"
	"github.com/kuma0328/music_repo/domain/repositry"
)

var _IMusicUsecase = &MusicUsecase{}

type IMusicUsecase interface {
	CreateMusic(*entity.MusicEntity) error
}

type MusicUsecase struct {
	repo repositry.IMusicRepositry
}

func NewMusicUsecase(repo repositry.IMusicRepositry) *MusicUsecase {
	return &MusicUsecase{
		repo: repo,
	}
}

func (mu *MusicUsecase) CreateMusic(music *entity.MusicEntity) error {
	if music.ID == 0 {
		return fmt.Errorf("ID is empty")
	}
	
	if !entity.IsYoutubeURL(music.URL) {
		return fmt.Errorf("URL is invalid")
	}

	return mu.repo.CreateMusic(music)
}

