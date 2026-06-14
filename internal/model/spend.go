package model

import "time"

type SpendRequest struct {
	CampaignID   int64 `json:"campaign_id"`
	AmountMicros int64 `json:"amount_micros"`
}

type SpendResponse struct {
	ID                 string `json:"id"`
	RunningTotalMicros int64  `json:"running_total_micros"`
}

type SpendEvent struct {
	ID           string
	TenantID     int64
	CampaignID   int64
	AmountMicros int64
	CreatedAt    time.Time
}
