package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type MealHandler struct {
	logger  *prettylog.Logger
	service service.MealService
}

func NewMealHandler(service service.MealService, logger *prettylog.Logger) *MealHandler {
	return &MealHandler{logger: logger, service: service}
}

func (h *MealHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	req, err := decode[model.Meal](r)
	if err != nil {
		h.logger.Errorf("Error bad create meal request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Date.IsZero() {
		req.Date = time.Now()
	}

	user, err := getUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	req.User = user

	meal, err := h.service.CreateMeal(r.Context(), req)
	if err != nil {
		h.logger.Errorf("Error creating meal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Created meal successfully with ID=%d", meal.Id)
	err = encode(w, r, 201, meal)
	if err != nil {
		h.logger.Errorf("Error encoding created meal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *MealHandler) GetMealById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	meal, err := h.service.GetMealById(r.Context(), intId)
	if err != nil {
		h.logger.Errorf("Error fetching meal with id=%d: %s", intId, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 200, meal)
	if err != nil {
		h.logger.Errorf("Error encoding meal in GetMealById: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *MealHandler) GetMealsFromDateRange(w http.ResponseWriter, r *http.Request) {
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

	meals, err := h.service.GetMealsFromDateRange(r.Context(), user, dates["from"], dates["to"])
	if err != nil {
		h.logger.Errorf("Error fetching meals from range: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 200, meals)
	if err != nil {
		h.logger.Errorf("Error encoding response in GetMealsFromDateRange: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
