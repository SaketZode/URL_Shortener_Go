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
	urlService service.UrlService
)

func init() {
	// initialize data store
	dataStore = store.NewDatabase()
	// initialize URL service and inject the data store
	urlService = &service.UrlServiceStruct{Store: dataStore}
}

// API for generating the short(encoded) URL
func CreateShortUrl(c *gin.Context) {
	// parse request body
	request := models.EncodingRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	// encoding the URL to short URL
	encodedUrl := urlService.EncodeURL(request)

	// send response
	c.JSON(http.StatusOK, encodedUrl)
}

// API to retrive the original URL form the encoded(short) URL
func DecodeShortUrl(c *gin.Context) {
	// parsing request body
	request := models.DecodingRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Message: err.Error(),
		})
		return
	}
	// decoding the short URL
	decodedResponse, decodingErr := urlService.DecodeURL(request)
	if decodingErr != nil {
		c.JSON(http.StatusInternalServerError, decodingErr)
		return
	}
	// send response
	c.JSON(http.StatusOK, decodedResponse)
}
