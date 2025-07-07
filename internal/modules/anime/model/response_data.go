package model

import "github.com/google/uuid"

type AnimeData struct {
	ID     uuid.UUID `json:"id"`
	Title  string    `json:"title"`
	Studio string    `json:"studio"`
	Years  string    `json:"years"`
	Season string    `json:"season"`
}
