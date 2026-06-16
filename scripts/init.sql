CREATE TABLE IF NOT EXISTS spend_events (
    id TEXT PRIMARY KEY,
    tenant_id BIGINT NOT NULL,
    campaign_id BIGINT NOT NULL,
    amount_micros BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);