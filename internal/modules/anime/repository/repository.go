package anime

import (
	"context"

	"gorm.io/gorm"
)

type animeRepository struct {
	db *gorm.DB
}

func (animerepository *animeRepository) Create(ctx context.Context, anime *Anime) error {
	panic("not implemented") // TODO: Implement
}

func (animerepository *animeRepository) Update(ctx context.Context, anime *Anime) error {
	panic("not implemented") // TODO: Implement
}

func (animerepository *animeRepository) Delete(ctx context.Context, id string) error {
	panic("not implemented") // TODO: Implement
}

func (animerepository *animeRepository) FindAll(ctx context.Context) ([]Anime, error) {
	panic("not implemented") // TODO: Implement
}

func (animerepository *animeRepository) FindById(ctx context.Context, id string) (Anime, error) {
	panic("not implemented") // TODO: Implement
}

type AnimeRepository interface {
	Create(ctx context.Context, anime *Anime) error
	Update(ctx context.Context, anime *Anime) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context) ([]Anime, error)
	FindById(ctx context.Context, id string) (Anime, error)
}
