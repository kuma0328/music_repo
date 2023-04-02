package entity

import (
	"regexp"
	"time"
)

type MusicEntity struct {
	ID 				int
	URL 			string
	Comment 	string
	ImageData []byte
	PostCount int
	CreatedAt time.Time
}

type PlayListEntity []*MusicEntity

func IsYoutubeURL(url string) bool {
	pattern := `^(https?\:\/\/)?(www\.)?(youtube\.com|youtu\.be)\/.+`
	r, err := regexp.Compile(pattern)
	if err != nil {
			// 正規表現のコンパイルに失敗した場合は、falseを返す
			return false
	}
	return r.MatchString(url)
}