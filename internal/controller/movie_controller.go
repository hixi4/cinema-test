package controller

import (
	"cinema-test/internal/model"
	"cinema-test/internal/service"
	"encoding/json"
	"net/http"
)

// MovieController надає контроллер для фильмів
type MovieController struct {
	service *service.MovieService
}

// NewMovieController створює новий контроллер фільмів
func NewMovieController(s *service.MovieService) *MovieController {
	return &MovieController{service: s}
}

// ListMovies вертає список усіх фільмів
func (c *MovieController) ListMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := c.service.GetMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(movies)
}

// OrderMovie обробляє замовлення білета на фільм
func (c *MovieController) OrderMovie(w http.ResponseWriter, r *http.Request) {
	var order model.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	savedOrder, err := c.service.OrderMovie(order.MovieID, order.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(savedOrder)
}

// ListOrders повертає список усіх замовлень
func (c *MovieController) ListOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := c.service.GetOrders()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(orders)
}
