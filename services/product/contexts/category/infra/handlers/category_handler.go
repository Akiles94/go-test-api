package handlers

import (
	"net/http"
	"strconv"

	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/dto"
	"github.com/Akiles94/go-test-api/services/product/contexts/category/application/ports/inbound"
	"github.com/Akiles94/go-test-api/shared/infra/shared_handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const topLimitValue = 100
const bottomLimitValue = 1

// CategoryHandler handles HTTP requests for categories
type CategoryHandler struct {
	createCategoryUseCase   inbound.CreateCategoryUseCasePort
	updateCategoryUseCase   inbound.UpdateCategoryUseCasePort
	patchCategoryUseCase    inbound.PatchCategoryUseCasePort
	deleteCategoryUseCase   inbound.DeleteCategoryUseCasePort
	getAllCategoriesUseCase inbound.GetAllCategoriesUseCasePort
	getOneCategoryUseCase   inbound.GetOneCategoryUseCasePort
}

// NewCategoryHandler creates a new CategoryHandler
func NewCategoryHandler(
	createCategoryUseCase inbound.CreateCategoryUseCasePort,
	updateCategoryUseCase inbound.UpdateCategoryUseCasePort,
	patchCategoryUseCase inbound.PatchCategoryUseCasePort,
	deleteCategoryUseCase inbound.DeleteCategoryUseCasePort,
	getAllCategoriesUseCase inbound.GetAllCategoriesUseCasePort,
	getOneCategoryUseCase inbound.GetOneCategoryUseCasePort,
) *CategoryHandler {
	return &CategoryHandler{
		createCategoryUseCase:   createCategoryUseCase,
		updateCategoryUseCase:   updateCategoryUseCase,
		patchCategoryUseCase:    patchCategoryUseCase,
		deleteCategoryUseCase:   deleteCategoryUseCase,
		getAllCategoriesUseCase: getAllCategoriesUseCase,
		getOneCategoryUseCase:   getOneCategoryUseCase,
	}
}

// GetPaginated godoc
// @Summary Get paginated categories
// @Description Get a paginated list of categories with optional cursor and limit
// @Tags categories
// @Accept json
// @Produce json
// @Param cursor query string false "Cursor for pagination (UUID)"
// @Param limit query int false "Limit of categories per page (1-100)" minimum(1) maximum(100)
// @Success 200 {object} dto.PaginatedCategoryResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories [get]
func (ch *CategoryHandler) GetPaginated(c *gin.Context) {
	cursor := c.Query("cursor")
	limitStr := c.Query("limit")

	limit := 10 // default limit
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil || parsedLimit < bottomLimitValue || parsedLimit > topLimitValue {
			c.Error(shared_handlers.ErrInvalidLimit)
			return
		}
		limit = parsedLimit
	}

	var cursorPtr *string
	if cursor != "" {
		if _, err := uuid.Parse(cursor); err != nil {
			c.Error(shared_handlers.ErrInvalidCursor)
			return
		}
		cursorPtr = &cursor
	}

	result, err := ch.getAllCategoriesUseCase.Execute(c.Request.Context(), cursorPtr, limit)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetByID godoc
// @Summary Get category by ID
// @Description Get a single category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID (UUID)" format(uuid)
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories/{id} [get]
func (ch *CategoryHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	result, err := ch.getOneCategoryUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Create godoc
// @Summary Create a new category
// @Description Create a new category with the provided data
// @Tags categories
// @Accept json
// @Produce json
// @Param category body dto.CreateCategoryRequest true "Category data"
// @Success 201 {object} dto.CategoryResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories [post]
func (ch *CategoryHandler) Create(c *gin.Context) {
	var request dto.CreateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}

	result, err := ch.createCategoryUseCase.Execute(c.Request.Context(), request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, result)
}

// Update godoc
// @Summary Update a category
// @Description Update an existing category with the provided data
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID (UUID)" format(uuid)
// @Param category body dto.UpdateCategoryRequest true "Category data"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories/{id} [put]
func (ch *CategoryHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	var request dto.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}

	result, err := ch.updateCategoryUseCase.Execute(c.Request.Context(), id, request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Patch godoc
// @Summary Partially update a category
// @Description Partially update an existing category with the provided data
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID (UUID)" format(uuid)
// @Param category body dto.PatchCategoryRequest true "Category data"
// @Success 200 {object} dto.CategoryResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories/{id} [patch]
func (ch *CategoryHandler) Patch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	var request dto.PatchCategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}

	result, err := ch.patchCategoryUseCase.Execute(c.Request.Context(), id, request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, result)
}

// Delete godoc
// @Summary Delete a category
// @Description Delete an existing category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path string true "Category ID (UUID)" format(uuid)
// @Success 204 "No Content"
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /categories/{id} [delete]
func (ch *CategoryHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	err = ch.deleteCategoryUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
