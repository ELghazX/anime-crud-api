package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/elghazx/anime-crud-api/internal/common/dto"
	"github.com/elghazx/anime-crud-api/internal/modules/anime/model"
	"github.com/elghazx/anime-crud-api/internal/modules/anime/service"
	"github.com/gofiber/fiber/v2"
)

type AnimeHandler struct {
	service service.AnimeService
}

func NewAnimeHandler(service service.AnimeService) *AnimeHandler {
	return &AnimeHandler{service: service}
}

func (h *AnimeHandler) CreateAnime(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req model.CreateAnimeRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	err := h.service.Create(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusCreated).JSON(dto.CreateResponseSuccess(""))
}

func (h *AnimeHandler) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	res, err := h.service.Index(c)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponseError(err.Error()))
	}
	return ctx.Status(http.StatusOK).
		JSON(dto.CreateResponseSuccess(res))
}
