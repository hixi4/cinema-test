package service

import (
	"cinema/internal/model"
)

// MovieRepository інтерфейс для доступу до даних фільмів
type MovieRepository interface {
	GetAllMovies() ([]model.Movie, error)
	GetMovieByID(id int) (*model.Movie, error)
	SaveOrder(order *model.Order) (*model.Order, error)
	GetAllOrders() ([]model.Order, error)
}

// MovieService представляє служебний шар для роботи з фільмами
type MovieService struct {
	repo MovieRepository
}

// NewMovieService створює нову службу для фільмів
func NewMovieService(repo MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

// GetMovies повертає список усіх фільмів
func (s *MovieService) GetMovies() ([]model.Movie, error) {
	return s.repo.GetAllMovies()
}

// OrderMovie створює замовлення на фільм
func (s *MovieService) OrderMovie(movieID, userID int) (*model.Order, error) {
	order := &model.Order{
		MovieID: movieID,
		UserID:  userID,
	}
	return s.repo.SaveOrder(order)
}

// GetOrders повертає список усіх замовлень
func (s *MovieService) GetOrders() ([]model.Order, error) {
	return s.repo.GetAllOrders()
}
