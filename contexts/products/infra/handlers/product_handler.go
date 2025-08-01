package handlers

import (
	"net/http"
	"strconv"

	"github.com/Akiles94/go-test-api/contexts/products/application/dto"
	"github.com/Akiles94/go-test-api/contexts/products/application/ports/inbound"
	"github.com/Akiles94/go-test-api/contexts/shared/application/shared_dto"
	"github.com/Akiles94/go-test-api/contexts/shared/infra/shared_handlers"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const topLimitValue = 100
const bottomLimitValue = 1

// ProductHandler handles HTTP requests for products
type ProductHandler struct {
	createProductUseCase  inbound.CreateProductUseCasePort
	updateProductUseCase  inbound.UpdateProductUseCasePort
	patchProductUseCase   inbound.PatchProductUseCasePort
	deleteProductUseCase  inbound.DeleteProductUseCasePort
	getAllProductsUseCase inbound.GetAllProductsUseCasePort
	getOneProductUseCase  inbound.GetOneProductUseCasePort
}

// NewProductHandler creates a new ProductHandler
func NewProductHandler(createProductUseCase inbound.CreateProductUseCasePort, updateProductUseCase inbound.UpdateProductUseCasePort, patchProductUseCase inbound.PatchProductUseCasePort, deleteProductUseCase inbound.DeleteProductUseCasePort, getAllProductsUseCase inbound.GetAllProductsUseCasePort, getOneProductUseCase inbound.GetOneProductUseCasePort) *ProductHandler {
	return &ProductHandler{
		createProductUseCase:  createProductUseCase,
		updateProductUseCase:  updateProductUseCase,
		patchProductUseCase:   patchProductUseCase,
		deleteProductUseCase:  deleteProductUseCase,
		getAllProductsUseCase: getAllProductsUseCase,
		getOneProductUseCase:  getOneProductUseCase,
	}
}

// GetPaginated godoc
// @Summary Get paginated products
// @Description Get a paginated list of products with optional cursor and limit
// @Tags products
// @Accept json
// @Produce json
// @Param cursor query string false "Cursor for pagination (UUID)"
// @Param limit query int false "Limit of products per page (1-100)" minimum(1) maximum(100)
// @Success 200 {object} shared_dto.PaginatedResult[dto.ProductResponse]
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products [get]
func (ph *ProductHandler) GetPaginated(c *gin.Context) {
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")
	var limit *int
	var cursor *string
	if cursorStr != "" {
		_, err := uuid.Parse(cursorStr)
		if err != nil {
			c.Error(shared_handlers.ErrInvalidCursor)
			return
		}
		cursor = &cursorStr
	}
	if limitStr != "" {
		limitValue, err := strconv.Atoi(limitStr)
		if err != nil {
			c.Error(shared_handlers.ErrInvalidLimit)
			return
		}

		if limitValue < bottomLimitValue {
			c.Error(shared_handlers.ErrInvalidLimit)
			return
		}

		const maxLimit = topLimitValue
		if limitValue > maxLimit {
			c.Error(shared_handlers.ErrInvalidLimit)
			return
		}

		limit = &limitValue
	}
	products, nextCursor, err := ph.getAllProductsUseCase.Execute(c.Request.Context(), cursor, limit)
	if err != nil {
		c.Error(err)
		return
	}
	var productResponses []dto.ProductResponse = make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		productResponse := dto.NewProductResponseFromDomainModel(product)
		productResponses = append(productResponses, productResponse)
	}
	response := shared_dto.NewPaginatedResult(productResponses, nextCursor)

	c.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary Get product by ID
// @Description Get a single product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)" format(uuid)
// @Success 200 {object} dto.ProductResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products/{id} [get]
func (ph *ProductHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}
	product, err := ph.getOneProductUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		return
	}
	if product == nil {
		c.Error(shared_handlers.ErrNotFound)
		return
	}
	productResponse := dto.NewProductResponseFromDomainModel(product)
	c.JSON(http.StatusOK, productResponse)
}

// Create godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product creation details"
// @Success 201 {object} dto.ProductResponse
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products [post]
func (ph *ProductHandler) Create(c *gin.Context) {
	var productDto dto.CreateProductRequest
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}
	product, err := productDto.ToDomainModel()
	if err != nil {
		c.Error(err)
		return
	}

	if err := ph.createProductUseCase.Execute(c.Request.Context(), product); err != nil {
		c.Error(err)
		return
	}
	productResponse := dto.ProductResponse{
		ID:       product.ID(),
		Sku:      product.Sku(),
		Name:     product.Name(),
		Category: product.Category(),
		Price:    product.Price(),
	}

	c.JSON(http.StatusCreated, productResponse)
}

// Update godoc
// @Summary Update a product
// @Description Update all fields of a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)" format(uuid)
// @Param product body dto.CreateProductRequest true "Product update details"
// @Success 204
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products/{id} [put]
func (ph *ProductHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	var productDto dto.CreateProductRequest
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}
	product, err := productDto.ToDomainModel()
	if err != nil {
		c.Error(err)
		return
	}

	if err := ph.updateProductUseCase.Execute(c.Request.Context(), id, product); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Patch godoc
// @Summary Partially update a product
// @Description Update specific fields of a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)" format(uuid)
// @Param product body dto.PatchProductRequest true "Product patch details"
// @Success 204
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products/{id} [patch]
func (ph *ProductHandler) Patch(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	var patch dto.PatchProductRequest
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.Error(shared_handlers.ErrInvalidPayload)
		return
	}

	updates := make(map[string]interface{})
	if patch.Sku != nil {
		updates["sku"] = *patch.Sku
	}
	if patch.Name != nil {
		updates["name"] = *patch.Name
	}
	if patch.Category != nil {
		updates["category"] = *patch.Category
	}
	if patch.Price != nil {
		updates["price"] = decimal.NewFromFloat(*patch.Price)
	}

	if err := ph.patchProductUseCase.Execute(c.Request.Context(), id, updates); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}

// Delete godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID (UUID)" format(uuid)
// @Success 204
// @Failure 400 {object} shared_dto.ErrorResponse
// @Failure 404 {object} shared_dto.ErrorResponse
// @Failure 500 {object} shared_dto.ErrorResponse
// @Router /products/{id} [delete]
func (ph *ProductHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.Error(shared_handlers.ErrInvalidUUID)
		return
	}

	if err := ph.deleteProductUseCase.Execute(c.Request.Context(), id); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusNoContent)
}
