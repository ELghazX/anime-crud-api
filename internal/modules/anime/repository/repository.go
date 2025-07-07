package repository

import (
	"context"

	domain "github.com/elghazx/anime-crud-api/internal/modules/anime/model"
	"gorm.io/gorm"
)

type AnimeRepository interface {
	Create(ctx context.Context, anime *domain.Anime) error
	Update(ctx context.Context, anime *domain.Anime) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]domain.Anime, error)
	FindById(ctx context.Context, id string) (domain.Anime, error)
}

type animeRepository struct {
	db *gorm.DB
}

func NewAnimeRepository(db *gorm.DB) AnimeRepository {
	return &animeRepository{
		db: db,
	}
}

func (animerepository *animeRepository) Create(ctx context.Context, anime *domain.Anime) error {
	// panic("not implemented") // TODO: Implement
	if err := animerepository.db.WithContext(ctx).Create(anime).Error; err != nil {
		return err
	}
	return nil
}

func (animerepository *animeRepository) Update(ctx context.Context, anime *domain.Anime) error {
	panic("not implemented") // TODO: Implement
}

func (animerepository *animeRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

func (r *animeRepository) FindAll(ctx context.Context) (animes []domain.Anime, err error) {
	if err := r.db.WithContext(ctx).Find(&animes).Error; err != nil {
		return nil, err
	}
	return
}

func (animerepository *animeRepository) FindById(ctx context.Context, id string) (domain.Anime, error) {
	panic("not implemented") // TODO: Implement
}
