package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Akiles94/go-test-api/contexts/product/application/dto"
	"github.com/Akiles94/go-test-api/contexts/product/domain/models"
	"github.com/Akiles94/go-test-api/contexts/product/domain/models/models_mothers"
	"github.com/Akiles94/go-test-api/contexts/product/infra/handlers"
	"github.com/Akiles94/go-test-api/contexts/product/infra/handlers/handlers_mocks"
	"github.com/Akiles94/go-test-api/contexts/shared/application/shared_dto"
	"github.com/Akiles94/go-test-api/contexts/shared/infra/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ProductHandlerTestSuite struct {
	suite.Suite
	handler *handlers.ProductHandler
	router  *gin.Engine
	// Mocks - solo para poder hacer assertions y configurar expectations
	mockCreateUseCase *handlers_mocks.MockCreateProductUseCase
	mockUpdateUseCase *handlers_mocks.MockUpdateProductUseCase
	mockPatchUseCase  *handlers_mocks.MockPatchProductUseCase
	mockDeleteUseCase *handlers_mocks.MockDeleteProductUseCase
	mockGetAllUseCase *handlers_mocks.MockGetAllProductsUseCase
	mockGetOneUseCase *handlers_mocks.MockGetOneProductUseCase
}

func (suite *ProductHandlerTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
}

func (suite *ProductHandlerTestSuite) SetupTest() {
	// Create mocks
	suite.mockCreateUseCase = new(handlers_mocks.MockCreateProductUseCase)
	suite.mockUpdateUseCase = new(handlers_mocks.MockUpdateProductUseCase)
	suite.mockPatchUseCase = new(handlers_mocks.MockPatchProductUseCase)
	suite.mockDeleteUseCase = new(handlers_mocks.MockDeleteProductUseCase)
	suite.mockGetAllUseCase = new(handlers_mocks.MockGetAllProductsUseCase)
	suite.mockGetOneUseCase = new(handlers_mocks.MockGetOneProductUseCase)

	// Create handler with injected mocks
	suite.handler = handlers.NewProductHandler(
		suite.mockCreateUseCase,
		suite.mockUpdateUseCase,
		suite.mockPatchUseCase,
		suite.mockDeleteUseCase,
		suite.mockGetAllUseCase,
		suite.mockGetOneUseCase,
	)

	// Setup router
	suite.router = gin.New()
	suite.router.Use(middlewares.ErrorHandlerMiddleware())

	suite.router.GET("/products", suite.handler.GetPaginated)
	suite.router.GET("/products/:id", suite.handler.GetByID)
	suite.router.POST("/products", suite.handler.Create)
	suite.router.PUT("/products/:id", suite.handler.Update)
	suite.router.PATCH("/products/:id", suite.handler.Patch)
	suite.router.DELETE("/products/:id", suite.handler.Delete)
}

func (suite *ProductHandlerTestSuite) TearDownTest() {
	// Reset mocks after each test
	suite.mockCreateUseCase.ExpectedCalls = nil
	suite.mockUpdateUseCase.ExpectedCalls = nil
	suite.mockPatchUseCase.ExpectedCalls = nil
	suite.mockDeleteUseCase.ExpectedCalls = nil
	suite.mockGetAllUseCase.ExpectedCalls = nil
	suite.mockGetOneUseCase.ExpectedCalls = nil
}

func (suite *ProductHandlerTestSuite) TestGetPaginated() {
	suite.Run("should return paginated products successfully", func() {
		// Arrange
		products := []models.Product{
			models_mothers.NewProductMother().WithName("Product 1").MustBuild(),
			models_mothers.NewProductMother().WithName("Product 2").MustBuild(),
		}
		nextCursor := "next-cursor"

		suite.mockGetAllUseCase.On("Execute", mock.Anything, (*string)(nil), (*int)(nil)).
			Return(products, &nextCursor, nil)

		// Act
		req := httptest.NewRequest(http.MethodGet, "/products", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusOK, w.Code)

		var response shared_dto.PaginatedResult[dto.ProductResponse]
		err := json.Unmarshal(w.Body.Bytes(), &response)
		suite.NoError(err)
		suite.Len(response.Items, 2)
		suite.Equal(&nextCursor, response.NextCursor)

		suite.mockGetAllUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should handle pagination with limit", func() {
		// Arrange
		products := []models.Product{
			models_mothers.NewProductMother().WithName("Product 1").MustBuild(),
		}
		limit := 1

		suite.mockGetAllUseCase.On("Execute", mock.Anything, (*string)(nil), &limit).
			Return(products, (*string)(nil), nil)

		// Act
		req := httptest.NewRequest(http.MethodGet, "/products?limit=1", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusOK, w.Code)
		suite.mockGetAllUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return error for invalid cursor", func() {
		// Act
		req := httptest.NewRequest(http.MethodGet, "/products?cursor=invalid-uuid", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})

	suite.Run("should return error for invalid limit", func() {
		// Act
		req := httptest.NewRequest(http.MethodGet, "/products?limit=invalid", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func (suite *ProductHandlerTestSuite) TestGetByID() {
	suite.Run("should return product successfully", func() {
		// Arrange
		productID := uuid.New()
		product := models_mothers.NewProductMother().
			WithID(productID).
			WithName("Test Product").
			MustBuild()

		suite.mockGetOneUseCase.On("Execute", mock.Anything, productID).
			Return(product, nil)

		// Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", productID), nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusOK, w.Code)

		var response dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		suite.NoError(err)
		suite.Equal(productID, response.ID)
		suite.Equal("Test Product", response.Name)

		suite.mockGetOneUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return 404 when product not found", func() {
		// Arrange
		productID := uuid.New()

		suite.mockGetOneUseCase.On("Execute", mock.Anything, productID).
			Return(nil, nil)

		// Act
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/products/%s", productID), nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusNotFound, w.Code)
		suite.mockGetOneUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return error for invalid UUID", func() {
		// Act
		req := httptest.NewRequest(http.MethodGet, "/products/invalid-uuid", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func (suite *ProductHandlerTestSuite) TestCreate() {
	suite.Run("should create product successfully", func() {
		// Arrange
		createRequest := dto.CreateProductRequest{
			Sku:      "TEST-001",
			Name:     "Test Product",
			Category: "Test Category",
			Price:    99.99,
		}

		suite.mockCreateUseCase.On("Execute", mock.Anything, mock.AnythingOfType("*models.product")).
			Return(nil)

		// Act
		body, _ := json.Marshal(createRequest)
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusCreated, w.Code)

		var response dto.ProductResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		suite.NoError(err)
		suite.Equal(createRequest.Sku, response.Sku)
		suite.Equal(createRequest.Name, response.Name)

		suite.mockCreateUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return error for invalid payload", func() {
		// Act
		req := httptest.NewRequest(http.MethodPost, "/products", bytes.NewBuffer([]byte("invalid json")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func (suite *ProductHandlerTestSuite) TestUpdate() {
	suite.Run("should update product successfully", func() {
		// Arrange
		productID := uuid.New()
		updateRequest := dto.CreateProductRequest{
			Sku:      "UPDATED-001",
			Name:     "Updated Product",
			Category: "Updated Category",
			Price:    199.99,
		}

		suite.mockUpdateUseCase.On("Execute", mock.Anything, productID, mock.AnythingOfType("*models.product")).
			Return(nil)

		// Act
		body, _ := json.Marshal(updateRequest)
		req := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/products/%s", productID), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusNoContent, w.Code)
		suite.mockUpdateUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return error for invalid UUID", func() {
		// Act
		req := httptest.NewRequest(http.MethodPut, "/products/invalid-uuid", bytes.NewBuffer([]byte("{}")))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func (suite *ProductHandlerTestSuite) TestPatch() {
	suite.Run("should patch product successfully", func() {
		// Arrange
		productID := uuid.New()
		name := "Patched Name"
		price := 299.99
		patchRequest := dto.PatchProductRequest{
			Name:  &name,
			Price: &price,
		}

		expectedUpdates := map[string]interface{}{
			"name":  name,
			"price": decimal.NewFromFloat(price),
		}

		suite.mockPatchUseCase.On("Execute", mock.Anything, productID, expectedUpdates).
			Return(nil)

		// Act
		body, _ := json.Marshal(patchRequest)
		req := httptest.NewRequest(http.MethodPatch, fmt.Sprintf("/products/%s", productID), bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusNoContent, w.Code)
		suite.mockPatchUseCase.AssertExpectations(suite.T())
	})
}

func (suite *ProductHandlerTestSuite) TestDelete() {
	suite.Run("should delete product successfully", func() {
		// Arrange
		productID := uuid.New()

		suite.mockDeleteUseCase.On("Execute", mock.Anything, productID).
			Return(nil)

		// Act
		req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/products/%s", productID), nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusNoContent, w.Code)
		suite.mockDeleteUseCase.AssertExpectations(suite.T())
	})

	suite.Run("should return error for invalid UUID", func() {
		// Act
		req := httptest.NewRequest(http.MethodDelete, "/products/invalid-uuid", nil)
		w := httptest.NewRecorder()
		suite.router.ServeHTTP(w, req)

		// Assert
		suite.Equal(http.StatusBadRequest, w.Code)
	})
}

func TestProductHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(ProductHandlerTestSuite))
}
