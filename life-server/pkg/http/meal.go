package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinbather/life/life-server/pkg/model"
	"github.com/justinbather/life/life-server/pkg/service"
	"github.com/justinbather/prettylog"
)

type mealHandler struct {
	logger  *prettylog.Logger
	service service.MealService
}

func NewMealHandler(service service.MealService, logger *prettylog.Logger) *mealHandler {
	return &mealHandler{logger: logger, service: service}
}

func (h *mealHandler) CreateMeal(w http.ResponseWriter, r *http.Request) {
	req, err := decode[model.Meal](r)
	if err != nil {
		h.logger.Errorf("Error bad create meal request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	meal, err := h.service.CreateMeal(r.Context(), req)
	if err != nil {
		h.logger.Errorf("Error creating meal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = encode(w, r, 201, meal)
	if err != nil {
		h.logger.Errorf("Error encoding created meal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *mealHandler) GetMealById(w http.ResponseWriter, r *http.Request) {
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

func (h *mealHandler) GetMealsFromDateRange(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	from := params["from"]
	to := params["to"]

	if from == "" || to == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	fromDate, err := time.Parse(time.RFC3339, from)
	if err != nil {
		h.logger.Errorf("Error parsing from date. Given: %s. Threw: %s", from, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Date must be in YYYY-MM-DDTHH:MM:SST format"))
		return
	}

	toDate, err := time.Parse(time.RFC3339, to)
	if err != nil {
		h.logger.Errorf("Error parsing to date. Given: %s. Threw: %s", to, err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Date must be in YYYY-MM-DDTHH:MM:SST format"))
		return
	}

	meals, err := h.service.GetMealsFromDateRange(r.Context(), fromDate, toDate)
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
