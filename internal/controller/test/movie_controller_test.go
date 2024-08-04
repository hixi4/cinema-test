package controller_test

import (
	"bytes"
	"cinema-test/internal/controller"
	"cinema-test/internal/model"
	"cinema-test/internal/repository"
	"cinema-test/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

// TestListMovies тестує ендпоінт для отримання списку фільмів
func TestListMovies(t *testing.T) {
	repo := repository.NewMovieRepository()
	svc := service.NewMovieService(repo)
	ctrl := controller.NewMovieController(svc)

	req, err := http.NewRequest("GET", "/movies", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/movies", ctrl.ListMovies)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}

// TestOrderMovie тестує ендпоінт для замовлення квитка
func TestOrderMovie(t *testing.T) {
	repo := repository.NewMovieRepository()
	svc := service.NewMovieService(repo)
	ctrl := controller.NewMovieController(svc)

	// Створюємо правильне тіло запиту
	order := model.Order{
		MovieID: 1,
		UserID:  1,
	}
	body, _ := json.Marshal(order)

	req, err := http.NewRequest("POST", "/order", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Post("/order", ctrl.OrderMovie)
	router.ServeHTTP(rr, req)

	// Оновіть ці перевірки відповідно до вашої логіки
	assert.Equal(t, http.StatusOK, rr.Code)
}

// TestListOrders тестує ендпоінт для отримання списку замовлень
func TestListOrders(t *testing.T) {
	repo := repository.NewMovieRepository()
	svc := service.NewMovieService(repo)
	ctrl := controller.NewMovieController(svc)

	req, err := http.NewRequest("GET", "/orders", nil)
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	router := chi.NewRouter()
	router.Get("/orders", ctrl.ListOrders)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
