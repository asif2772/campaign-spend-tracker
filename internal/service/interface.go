package service

import (
	"context"

	"github.com/asif2772/campaign-spend-tracker/internal/model"
)

type SpendService interface {
	CreateSpend(
		ctx context.Context,
		request model.SpendRequest,
	) (*model.SpendResponse, error)

	GetDailySpend(
		ctx context.Context,
		campaignID int64,
	) (int64, error)
}
