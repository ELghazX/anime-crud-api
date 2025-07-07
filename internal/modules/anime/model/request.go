package model

type CreateAnimeRequest struct {
	Title  string `json:"title" validate:"required"`
	Studio string `json:"studio" validate:"required"`
	Years  string `json:"years" validate:"required"`
	Season string `json:"season" validate:"required"`
}
