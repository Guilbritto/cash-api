package categories

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/mappers"
)

func (u *Service) GetAll(userId string) (*[]dto.CategoryResponse, error) {
	categories, err := u.CategoryRepository.GetAll(userId)
	if err != nil {
		return &[]dto.CategoryResponse{}, err
	}

	return mappers.CategoriesToResponse(*categories), nil
}
