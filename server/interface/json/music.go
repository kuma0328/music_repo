package json

import (
	"time"

	"github.com/kuma0328/music_repo/domain/entity"
)

type MusicJson struct {
	ID 				int 	    `json:"id"`
	URL 			string    `json:"url"`
	Comment 	string    `json:"comment"`
	ImageData []byte    `json:"image_data"`
	PostCount int       `json:"post_count"`
	CreatedAt time.Time `json:"created_at"`
}

func MusicEntityToJson(m *entity.MusicEntity) *MusicJson {
	return &MusicJson{
		ID: 			 m.ID,
		URL: 			 m.URL,
		Comment: 	 m.Comment,
		ImageData: m.ImageData,
		PostCount: m.PostCount,
		CreatedAt: m.CreatedAt,
	}
}

func MusicJsonToEntity(j *MusicJson) *entity.MusicEntity {
	return &entity.MusicEntity{
		ID: 			 j.ID,
		URL: 			 j.URL,
		Comment: 	 j.Comment,
		ImageData: j.ImageData,
		PostCount: j.PostCount,
		CreatedAt: j.CreatedAt,
	}
}