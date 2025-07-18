package handlers

import (
	"net/http"
	"strconv"

	"github.com/Akiles94/go-test-api/application/ports/inbound"
	"github.com/Akiles94/go-test-api/domain/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductHandler struct {
	useCase inbound.ProductUseCasePort
}

func NewProductHandler(useCase inbound.ProductUseCasePort) *ProductHandler {
	return &ProductHandler{
		useCase: useCase,
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
		}
		cursor = &cursorStr
	}
	if limitStr != "" {
		limitValue, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse limit"})
		}
		limit = &limitValue
	}
	productsResponse, err := ph.useCase.GetPaginated(cursor, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}
	c.JSON(http.StatusOK, productsResponse)
}

func (ph *ProductHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}
	product, err := ph.useCase.GetByID(id)
	if err != nil {
		print("Error getting product by ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get product"})
		return
	}
	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (ph *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := ph.useCase.Create(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (ph *ProductHandler) Update(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid UUID"})
		return
	}

	var payload models.Product
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := ph.useCase.Update(id, payload); err != nil {
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

	var patch models.ProductPatch
	if err := c.ShouldBindJSON(&patch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
		return
	}

	if err := ph.useCase.Patch(id, patch); err != nil {
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

	if err := ph.useCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}

	c.Status(http.StatusNoContent)
}
