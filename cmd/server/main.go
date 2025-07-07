package main

import (
	"github.com/elghazx/anime-crud-api/internal/common/config"
	"github.com/elghazx/anime-crud-api/internal/common/database"
	"github.com/elghazx/anime-crud-api/internal/modules/anime"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	dbConnection := database.GetDatabase(cnf.Database)

	app := fiber.New()

	animeModule := anime.NewModule(dbConnection)
	animeModule.RegisterRoutes(app)
	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
