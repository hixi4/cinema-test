package model

// Movie представляє структуру фільма
type Movie struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Order представляє замовлення на фільм
type Order struct {
	MovieID int
	UserID  int
}
