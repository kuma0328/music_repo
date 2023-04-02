package persistance

import (
	"fmt"
	"time"

	"github.com/kuma0328/music_repo/domain/entity"
	"github.com/kuma0328/music_repo/domain/repositry"
	"github.com/kuma0328/music_repo/infrastructure/database"
)

var _ repositry.IMusicRepositry = &MusicRepositry{}

type MusicRepositry struct {
	conn *database.Conn
}

func NewMusicRepositry(conn *database.Conn) repositry.IMusicRepositry {
	return &MusicRepositry {
		conn: conn,
	}
}

func (repo *MusicRepositry) CreateMusic(music *entity.MusicEntity) error {
	dto := MusicEntityToDto(music)
	query := `
	INSERT INTO music (id, url, comment, image_data, post_count, created_at)
	VALUES (:id, :url, :comment, :image_data, :post_count, :created_at)
	`

	_, err := repo.conn.DB.Exec(query, &dto)
	if err != nil {
		return fmt.Errorf("create music exec error : %w", err)
	}
	return nil
}

type musicDto struct {
	ID        int       `db:"id"`
	URL       string    `db:"url"`
	Comment   string    `db:"comment"`
	ImageData []byte    `db:"image_data"`
	PostCount int       `db:"post_count"`
	CreatedAt time.Time `db:"created_at"`
}

func MusicEntityToDto(m *entity.MusicEntity) musicDto {
	return musicDto{
		ID:        m.ID,
		URL:       m.URL,
		Comment:   m.Comment,
		ImageData: m.ImageData,
		PostCount: m.PostCount,
		CreatedAt: m.CreatedAt,
	}
}