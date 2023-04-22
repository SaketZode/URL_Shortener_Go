package handler

import (
	"net/http"
	"urlshortener/models"
	"urlshortener/service"
	"urlshortener/store"

	"github.com/gin-gonic/gin"
)

var (
	dataStore  store.DataStore
	urlService service.UrlEncodeService
)

func init() {
	dataStore = store.NewDatabase()
	urlService = &service.UrlEncoder{Store: dataStore}
}

func CreateShortUrl(c *gin.Context) {
	request := models.EncodingRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	encodedUrl, encodingErr := urlService.EncodeURL(request)
	if encodingErr != nil {
		c.JSON(http.StatusInternalServerError, encodingErr)
		return
	}
	c.JSON(http.StatusOK, encodedUrl)
}

func HandleShortUrlRedirect(c *gin.Context) {
}
