package handler

import (
	"encoding/json"
	"net/http"

	"github.com/kuma0328/music_repo/usecase"
)

type MusicHanlder struct {
	mu usecase.IMusicUsecase
}

func NewMusicHandler(mu usecase.IMusicUsecase) *MusicHanlder {
	return &MusicHanlder{
		mu : mu,
	}
}

func (h *MusicHanlder) CreateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	healthData := map[string]string{"health": "good!"}

	jsonResponse, err := json.Marshal(healthData)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Write(jsonResponse)
}