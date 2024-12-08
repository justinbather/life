package handlers

import (
	"net/http"
	"time"

	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type BmrHandler struct {
	logger  *prettylog.Logger
	service service.BmrService
}

func NewBmrHandler(service service.BmrService, logger *prettylog.Logger) *BmrHandler {
	return &BmrHandler{logger: logger, service: service}
}

func (h *BmrHandler) CreateBmr(w http.ResponseWriter, r *http.Request) {
	req, err := decode[model.Bmr](r)
	if err != nil {
		h.logger.Errorf("Error bad create bmr request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.CreatedAt.IsZero() {
		req.CreatedAt = time.Now()
	}

	user, err := getUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	req.UserId = user

	bmr, err := h.service.CreateBmr(r.Context(), req)
	if err != nil {
		h.logger.Errorf("Error creating bmr: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Created bmr successfully with ID=%d", bmr.Id)
	err = encode(w, r, 201, bmr)
	if err != nil {
		h.logger.Errorf("Error encoding created bmr: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *BmrHandler) GetMealById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (h *BmrHandler) GetBmrFromDateRange(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	dates, err := parseDateParams(r)
	if err != nil {
		h.logger.Errorf("Failed to parse date: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bmr, err := h.service.GetBmrsFromDateRange(r.Context(), user, dates["from"], dates["to"])
	if err != nil {
		h.logger.Errorf("Error fetching bmrs from range: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 200, bmr)
	if err != nil {
		h.logger.Errorf("Error encoding response in GetBmrFromDateRange: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
