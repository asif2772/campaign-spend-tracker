package repository

import (
	"context"
	"fmt"

	"github.com/asif2772/campaign-spend-tracker/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		pool: pool,
	}
}

func (r *PostgresRepository) InsertSpendEvent(
	ctx context.Context,
	event *model.SpendEvent,
) error {

	event.ID = uuid.NewString()

	query := `
	INSERT INTO spend_events
	(id, tenant_id, campaign_id, amount_micros)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.pool.Exec(
		ctx,
		query,
		event.ID,
		event.TenantID,
		event.CampaignID,
		event.AmountMicros,
	)

	if err != nil {
		return fmt.Errorf("insert spend event: %w", err)
	}

	return nil
}

func (r *PostgresRepository) Health(ctx context.Context) error {

	if err := r.pool.Ping(ctx); err != nil {
		return fmt.Errorf("postgres health check: %w", err)
	}

	return nil
}
