package handlers

import (
	"app/models"
	"app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetAll(c *gin.Context) {
	// Get pagination parameters from query string
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	// Get filters from query string
	filters := make(map[string]interface{})
	if name := c.DefaultQuery("name", ""); name != "" {
		filters["name"] = name
	}
	if price := c.DefaultQuery("price", ""); price != "" {
		priceVal, _ := strconv.ParseFloat(price, 64)
		filters["price"] = priceVal
	}

	// Get paginated and filtered product list
	products, total, err := h.Service.GetAllProducts(page, pageSize, filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	// Return paginated response
	c.JSON(http.StatusOK, gin.H{
		"data":      products,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	})
}

func (h *ProductHandler) Create(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	if err := h.Service.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.Service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}
	if err := h.Service.UpdateProduct(uint(id), &product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.Status(http.StatusNoContent)
}