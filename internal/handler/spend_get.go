package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/asif2772/campaign-spend-tracker/internal/tenant"
)

func (h *SpendHandler) GetDailySpend(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tenantHeader := r.Header.Get("X-Tenant-ID")

	tenantID, err := strconv.ParseInt(tenantHeader, 10, 64)
	if err != nil || tenantID <= 0 {
		http.Error(w, "invalid tenant id", http.StatusBadRequest)
		return
	}

	ctx := tenant.WithTenantID(r.Context(), tenantID)

	parts := strings.Split(r.URL.Path, "/")

	if len(parts) < 3 {
		http.Error(w, "invalid campaign id", http.StatusBadRequest)
		return
	}

	campaignID, err := strconv.ParseInt(parts[2], 10, 64)

	if err != nil {
		http.Error(w, "invalid campaign id", http.StatusBadRequest)
		return
	}

	total, err := h.service.GetDailySpend(
		ctx,
		campaignID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]any{
		"campaign_id":  campaignID,
		"total_micros": total,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(response)

}
