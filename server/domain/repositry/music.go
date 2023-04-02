package repositry

import "github.com/kuma0328/music_repo/domain/entity"

type IMusicRepositry interface {
	CreateMusic(*entity.MusicEntity) error
}