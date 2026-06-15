package handler

import (
	"net/http"

	"github.com/asif2772/campaign-spend-tracker/internal/repository"
)

type HealthHandler struct {
	postgres *repository.PostgresRepository
	redis    *repository.RedisRepository
}

func NewHealthHandler(
	postgres *repository.PostgresRepository,
	redis *repository.RedisRepository,
) *HealthHandler {
	return &HealthHandler{
		postgres: postgres,
		redis:    redis,
	}
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if err := h.postgres.Health(ctx); err != nil {
		http.Error(w, "postgres unavailable", http.StatusServiceUnavailable)
		return
	}

	if err := h.redis.Health(ctx); err != nil {
		http.Error(w, "redis unavailable", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("OK")); err != nil {
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
