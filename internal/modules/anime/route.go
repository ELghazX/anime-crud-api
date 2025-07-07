package anime

import (
	"github.com/elghazx/anime-crud-api/internal/modules/anime/handler"
	"github.com/elghazx/anime-crud-api/internal/modules/anime/repository"
	"github.com/elghazx/anime-crud-api/internal/modules/anime/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Module struct {
	handler *handler.AnimeHandler
}

func NewModule(db *gorm.DB) *Module {
	repo := repository.NewAnimeRepository(db)
	svc := service.NewAnimeService(repo)
	h := handler.NewAnimeHandler(svc)
	return &Module{handler: h}
}

func (m *Module) RegisterRoutes(app *fiber.App) {
	group := app.Group("/anime")
	group.Get("/", m.handler.Index)
	group.Post("/", m.handler.CreateAnime)
}
