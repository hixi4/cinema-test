package repository_test

import (
	"cinema-test/internal/model"
	"cinema-test/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGetAllMovies тестує метод отримання всіх фільмів
func TestGetAllMovies(t *testing.T) {
	repo := repository.NewMovieRepository()

	movies, err := repo.GetAllMovies()
	assert.NoError(t, err)
	assert.NotEmpty(t, movies)
}

// TestGetMovieByID тестує метод отримання фільму за ID
func TestGetMovieByID(t *testing.T) {
	repo := repository.NewMovieRepository()

	movie, err := repo.GetMovieByID(1)
	assert.NoError(t, err)
	assert.NotNil(t, movie)
}

// TestSaveOrder тестує метод збереження замовлення
func TestSaveOrder(t *testing.T) {
	repo := repository.NewMovieRepository()

	order := &model.Order{
		MovieID: 1,
		UserID:  1,
	}

	savedOrder, err := repo.SaveOrder(order)
	assert.NoError(t, err)
	assert.NotNil(t, savedOrder)
}
