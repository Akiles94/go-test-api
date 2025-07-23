package handlers

import (
	"net/http"
	"strconv"

	"github.com/Akiles94/go-test-api/products/application/dto"
	"github.com/Akiles94/go-test-api/products/application/ports/inbound"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductHandler struct {
	createProductUseCase  inbound.CreateProductUseCasePort
	updateProductUseCase  inbound.UpdateProductUseCasePort
	patchProductUseCase   inbound.PatchProductUseCasePort
	deleteProductUseCase  inbound.DeleteProductUseCasePort
	getAllProductsUseCase inbound.GetAllProductsUseCasePort
	getOneProductUseCase  inbound.GetOneProductUseCasePort
}

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

func (ph *ProductHandler) GetPaginated(c *gin.Context) {
	cursorStr := c.Query("cursor")
	limitStr := c.Query("limit")
	var limit *int
	var cursor *string
	if cursorStr != "" {
		_, err := uuid.Parse(cursorStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid cursor UUID"})
			return
		}
		cursor = &cursorStr
	}
	if limitStr != "" {
		limitValue, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse limit"})
			return
		}
		limit = &limitValue
	}
	products, nextCursor, err := ph.getAllProductsUseCase.Execute(cursor, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}
	var productResponses []dto.ProductResponse = make([]dto.ProductResponse, 0, len(products))
	for _, product := range products {
		productResponse := dto.ProductResponse{
			ID:       product.ID(),
			Sku:      product.Sku(),
			Name:     product.Name(),
			Category: product.Category(),
			Price:    product.Price(),
		}
		productResponses = append(productResponses, productResponse)
	}
	response := dto.ProductsListResponse{
		Products:   productResponses,
		NextCursor: nextCursor,
	}

	c.JSON(http.StatusOK, response)
}

func (ph *ProductHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}
	product, err := ph.getOneProductUseCase.Execute(id)
	if err != nil {
		print("Error getting product by ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get product"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	productResponse := dto.ProductResponse{
		ID:       product.ID(),
		Sku:      product.Sku(),
		Name:     product.Name(),
		Category: product.Category(),
		Price:    product.Price(),
	}
	c.JSON(http.StatusOK, productResponse)
}

func (ph *ProductHandler) Create(c *gin.Context) {
	var productDto dto.CreateProductRequest
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	product, err := productDto.ToDomainModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := ph.createProductUseCase.Execute(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
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

func (ph *ProductHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var productDto dto.CreateProductRequest
	if err := c.ShouldBindJSON(&productDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}
	product, err := productDto.ToDomainModel()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := ph.updateProductUseCase.Execute(id, product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (ph *ProductHandler) Patch(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var patch dto.PatchProductRequest
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
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

	if err := ph.patchProductUseCase.Execute(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to patch product"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (ph *ProductHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	if err := ph.deleteProductUseCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.Status(http.StatusNoContent)
}
