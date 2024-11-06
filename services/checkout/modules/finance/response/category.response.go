package response

import (
	"time"
)

type CreateCategoryResponse struct {
	ID        int       `json:"id"`
	UserID    *int      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateCategoryResponse struct {
	ID        int       `json:"id"`
	UserID    *int      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCategoryResponse struct {
	ID        int       `json:"id"`
	UserID    *int      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindCategoryResponse struct {
	ID        int       `json:"id"`
	UserID    *int      `json:"user_id"`
	Name      string    `json:"name"`
	Icon      string    `json:"icon"`
	Color     string    `json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FindAllCategoriesResponse struct {
	Items      []FindCategoryResponse `json:"items"`
	Total      int                    `json:"total"`
	TotalPages int                    `json:"total_pages"`
	Page       int                    `json:"page"`
	Limit      int                    `json:"limit"`
}
