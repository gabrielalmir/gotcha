package controllers

import (
	"net/http"

	"gotcha/src/services"

	"github.com/gin-gonic/gin"
)

type URLController struct {
	urlService *services.URLService
}

func NewURLController(urlService *services.URLService) *URLController {
	return &URLController{
		urlService: urlService,
	}
}

type CreateURLRequest struct {
	URL string `json:"url" binding:"required"`
}

func (c *URLController) CreateShortURL(ctx *gin.Context) {
	var req CreateURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url, err := c.urlService.CreateShortURL(ctx, req.URL)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	ctx.JSON(http.StatusCreated, url)
}

func (c *URLController) RedirectToOriginal(ctx *gin.Context) {
	shortURL := ctx.Param("short")

	url, err := c.urlService.GetByShortURL(ctx, shortURL)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

	ctx.Redirect(http.StatusMovedPermanently, url.Original)
}
