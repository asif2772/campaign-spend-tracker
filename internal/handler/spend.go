package handler

import "github.com/asif2772/campaign-spend-tracker/internal/service"

type SpendHandler struct {
	service service.SpendService
}

func NewSpendHandler(service service.SpendService) *SpendHandler {
	return &SpendHandler{
		service: service,
	}
}
