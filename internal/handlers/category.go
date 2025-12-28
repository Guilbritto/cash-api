package handlers

import (
	"github.com/Guilbritto/cash-api/internal/dto"
	"github.com/Guilbritto/cash-api/internal/models"
	"github.com/gofiber/fiber/v2"
)

type CategoryResponse = models.Category

func (h *Handlers) RegisterCategoriesEndpoint(api fiber.Router) {
	api.Post("/category", h.CreateCategory)
	api.Get("/category", h.ListCategories)
}

// CreateCategory godoc
// @Summary      Cria uma nova categoria
// @Description  Cria uma categoria para o usuário autenticado
// @Tags         categories
// @Accept       json
// @Produce      json
// @Param        body  body      dto.CreateCategoryRequest  true  "Dados da categoria"
// @Success      201   {object}  CategoryResponse
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Router       /categories [post]
func (h *Handlers) CreateCategory(c *fiber.Ctx) error {
	category := new(dto.CreateCategoryRequest)

	if err := c.BodyParser(category); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid category",
		})
	}

	createdCategory, err := h.UseCases.CategoriesUseCase.Create(*category, c.Locals("userId").(string))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create category",
		})
	}

	return c.JSON(createdCategory)
}

// CreateCategory godoc
// @Summary      Lista todas as categorias de um usuário ou globais
// @Description  Lista todas as categorias de um usuário ou globais
// @Tags         categories
// @Accept       json
// @Produce      json
// @Success      201   {object}  []models.Category
// @Failure      400   {object}  dto.ErrorResponse
// @Failure      401   {object}  dto.ErrorResponse
// @Router       /categories [post]
func (h *Handlers) ListCategories(c *fiber.Ctx) error {
	userId := c.Locals("userId").(string)
	categories, err := h.UseCases.CategoriesUseCase.GetAll(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server error",
		})
	}

	return c.JSON(categories)
}
