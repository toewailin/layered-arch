package handlers

import (
	"app/models"
	"app/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FaqHandler struct {
	Service *services.FaqService
}

func NewFaqHandler(service *services.FaqService) *FaqHandler {
	return &FaqHandler{Service: service}
}

func (h *FaqHandler) GetAll(c *gin.Context) {
	faqs, err := h.Service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch FAQs"})
		return
	}
	c.JSON(http.StatusOK, faqs)
}

func (h *FaqHandler) Create(c *gin.Context) {
	var faq models.Faq
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FAQ data"})
		return
	}
	if err := h.Service.Create(&faq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create FAQ"})
		return
	}
	c.JSON(http.StatusCreated, faq)
}

func (h *FaqHandler) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	faq, err := h.Service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ not found"})
		return
	}
	c.JSON(http.StatusOK, faq)
}

func (h *FaqHandler) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var faq models.Faq
	if err := c.ShouldBindJSON(&faq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid FAQ data"})
		return
	}
	if err := h.Service.Update(uint(id), &faq); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update FAQ"})
		return
	}
	c.JSON(http.StatusOK, faq)
}

func (h *FaqHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete FAQ"})
		return
	}
	c.Status(http.StatusNoContent)
}
