package repository

import (
	"fmt"
	"time"
)

func buildDailySpendKey(
	tenantID int64,
	campaignID int64,
	date time.Time,
) string {

	return fmt.Sprintf(
		"t:%d:daily_spend:%d:%s",
		tenantID,
		campaignID,
		date.Format("2006-01-02"),
	)
}
