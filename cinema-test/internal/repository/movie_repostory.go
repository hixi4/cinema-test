package repository

import (
	"cinema/internal/model"
	"errors"
)

// MovieRepository реалізує інтерфейс для доступу до даних фільмів
type MovieRepository struct {
	movies []model.Movie
	orders []model.Order
}

// NewMovieRepository створює новий репозиторій для фільмів
func NewMovieRepository() *MovieRepository {
	return &MovieRepository{
		movies: []model.Movie{
			{ID: 1, Title: "Movie 1", Description: "Description 1"},
			{ID: 2, Title: "Movie 2", Description: "Description 2"},
		},
		orders: []model.Order{},
	}
}

// GetAllMovies повертає список усіх фільмів
func (r *MovieRepository) GetAllMovies() ([]model.Movie, error) {
	return r.movies, nil
}

// GetMovieByID повертає фільм по ID
func (r *MovieRepository) GetMovieByID(id int) (*model.Movie, error) {
	for _, movie := range r.movies {
		if movie.ID == id {
			return &movie, nil
		}
	}
	return nil, errors.New("movie not found")
}

// SaveOrder зберігає замовлення
func (r *MovieRepository) SaveOrder(order *model.Order) (*model.Order, error) {
	r.orders = append(r.orders, *order)
	return order, nil
}

// GetAllOrders повертає список усіх замовлень
func (r *MovieRepository) GetAllOrders() ([]model.Order, error) {
	return r.orders, nil
}
