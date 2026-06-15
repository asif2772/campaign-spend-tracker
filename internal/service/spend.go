package service

import (
	"context"
	"fmt"
	"time"

	"github.com/asif2772/campaign-spend-tracker/internal/model"
	"github.com/asif2772/campaign-spend-tracker/internal/publisher"
	"github.com/asif2772/campaign-spend-tracker/internal/repository"
	"github.com/asif2772/campaign-spend-tracker/internal/tenant"
)

type spendService struct {
	postgres  *repository.PostgresRepository
	redis     *repository.RedisRepository
	publisher *publisher.Publisher
}

func NewSpendService(
	postgres *repository.PostgresRepository,
	redis *repository.RedisRepository,
	publisher *publisher.Publisher,
) SpendService {
	return &spendService{
		postgres:  postgres,
		redis:     redis,
		publisher: publisher,
	}
}
func (s *spendService) CreateSpend(
	ctx context.Context,
	request model.SpendRequest,
) (*model.SpendResponse, error) {

	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("tenant id not found in context")
	}

	event := &model.SpendEvent{
		TenantID:     tenantID,
		CampaignID:   request.CampaignID,
		AmountMicros: request.AmountMicros,
		CreatedAt:    time.Now().UTC(),
	}

	if err := s.postgres.InsertSpendEvent(ctx, event); err != nil {
		return nil, err
	}

	runningTotal, err := s.redis.IncrementDailySpend(
		ctx,
		tenantID,
		request.CampaignID,
		request.AmountMicros,
	)
	if err != nil {
		return nil, err
	}

	if err := s.publisher.Publish(ctx, *event); err != nil {
		return nil, err
	}

	return &model.SpendResponse{
		ID:                 event.ID,
		RunningTotalMicros: runningTotal,
	}, nil
}

func (s *spendService) GetDailySpend(
	ctx context.Context,
	campaignID int64,
) (int64, error) {

	tenantID, ok := tenant.FromContext(ctx)
	if !ok {
		return 0, fmt.Errorf("tenant id not found in context")
	}

	return s.redis.GetDailySpend(
		ctx,
		tenantID,
		campaignID,
	)
}
