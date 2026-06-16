package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/asif2772/campaign-spend-tracker/internal/model"
	"github.com/asif2772/campaign-spend-tracker/internal/service"
	"github.com/asif2772/campaign-spend-tracker/internal/tenant"
)

type SpendHandler struct {
	service service.SpendService
}

func NewSpendHandler(service service.SpendService) *SpendHandler {
	return &SpendHandler{
		service: service,
	}
}

func (h *SpendHandler) CreateSpend(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tenantHeader := r.Header.Get("X-Tenant-ID")
	if tenantHeader == "" {
		http.Error(w, "missing X-Tenant-ID", http.StatusBadRequest)
		return
	}

	tenantID, err := strconv.ParseInt(tenantHeader, 10, 64)
	if err != nil || tenantID <= 0 {
		http.Error(w, "invalid tenant id", http.StatusBadRequest)
		return
	}

	ctx := tenant.WithTenantID(r.Context(), tenantID)

	var request model.SpendRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if request.CampaignID <= 0 || request.AmountMicros <= 0 {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	response, err := h.service.CreateSpend(ctx, request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)
}
