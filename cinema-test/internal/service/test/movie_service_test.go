package service_test

import (
	"cinema/internal/repository"
	"cinema/internal/service"
	"testing"
)

func TestNewMovieService(t *testing.T) {
	repo := repository.NewMovieRepository()
	service := service.NewMovieService(repo)
	if service == nil {
		t.Error("Expected non-nil service")
	}
}
