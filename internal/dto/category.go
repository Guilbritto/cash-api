package dto

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryResponse struct {
	Id   string `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}
