package main

import (
	"cinema-test/internal/controller"
	"cinema-test/internal/repository"
	"cinema-test/internal/service"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Cinema API running")

	// Ініціалізація логгера
	logger := logrus.New()
	logger.Out = os.Stdout

	// Ініціалізація репозиторіїв
	repo := repository.NewMovieRepository()

	// Ініціалізація служб
	movieService := service.NewMovieService(repo)
	emailService := service.NewEmailService("smtp.example.com", "587", "username", "password")

	// Ініціалізація контролера
	movieController := controller.NewMovieController(movieService)

	// Ініціалізація роутера
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Роут для отримання списку доступних фільмів
	r.Get("/movies", movieController.ListMovies)

	// Роут для замовлення квитків
	r.Post("/order", movieController.OrderMovie)

	// Роут для отримання списку замовлених квитків
	r.Get("/orders", movieController.ListOrders)

	// Приклад використання EmailService
	err := emailService.SendEmail("recipient@example.com", "Test Subject", "Test Body")
	if err != nil {
		logger.Fatal("Failed to send email:", err)
	}

	// Запуск сервера на динамічному порту
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logger.Infof("Starting server on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		logger.Fatal("Failed to start server:", err)
	}
}
