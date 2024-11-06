package request

import "sendzap-checkout/services/checkout/modules/finance/filters"

type CreateCategoryRequest struct {
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required,min=3,max=50"`
	Icon   string `json:"icon" validate:"required,min=3,max=50"`
	Color  string `json:"color" validate:"required,min=3,max=50"`
}

type UpdateCategoryRequest struct {
	ID     int    `json:"id" validate:"required"`
	UserID int    `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required,min=3,max=50"`
	Icon   string `json:"icon" validate:"required,min=3,max=50"`
	Color  string `json:"color" validate:"required,min=3,max=50"`
}

type DeleteCategoryRequest struct {
	UserID int `json:"user_id" validate:"required"`
	ID     int `json:"id" validate:"required"`
}

type FindCategoryRequest struct {
	UserID int `json:"user_id" validate:"required"`
	ID     int `json:"id" validate:"required"`
}

type FindAllCategoriesRequest struct {
	UserID          int                     `json:"user_id" validate:"required"`
	CategoryFilters filters.CategoryFilters `json:"filters"`
}
