package handler

import (
	"github.com/iamseki/devices-api/src/repository"
)

type Handler struct {
	Repository *repository.Repository
}

func New(repo *repository.Repository) *Handler {
	return &Handler{repo}
}
