package service

import (
	"context"

	"github.com/elghazx/anime-crud-api/internal/modules/anime/model"
	"github.com/elghazx/anime-crud-api/internal/modules/anime/repository"
	"github.com/google/uuid"
)

type AnimeService interface {
	Create(ctx context.Context, req model.CreateAnimeRequest) error
	Index(ctx context.Context) ([]model.AnimeData, error)
}
type animeService struct {
	repo repository.AnimeRepository
}

func NewAnimeService(repo repository.AnimeRepository) AnimeService {
	return &animeService{
		repo: repo,
	}
}

func (animeservice *animeService) Create(ctx context.Context, req model.CreateAnimeRequest) error {
	anime := model.Anime{
		ID:     uuid.New(),
		Title:  req.Title,
		Studio: req.Studio,
		Years:  req.Years,
		Season: req.Season,
	}
	return animeservice.repo.Create(ctx, &anime)
}

func (s animeService) Index(ctx context.Context) ([]model.AnimeData, error) {
	animes, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var animeData []model.AnimeData
	for _, v := range animes {
		animeData = append(animeData, model.AnimeData{
			ID:     v.ID,
			Title:  v.Title,
			Studio: v.Studio,
			Years:  v.Years,
			Season: v.Season,
		})
	}
	return animeData, nil
}
