package handlers

import "rd/internal/db"

type Handler struct {
	DBLayer *db.Layer
}

func NewHandler(db *db.Layer) *Handler {
	return &Handler{
		DBLayer: db,
	}
}
