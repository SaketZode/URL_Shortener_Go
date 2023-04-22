package service

import (
	"urlshortener/constants"
	"urlshortener/hashgenerator"
	"urlshortener/models"
	"urlshortener/store"
)

// Setting(initializing) hash generator as MD5.
var hash hashgenerator.HashGenerator = &hashgenerator.MD5HashGenerator{}

type UrlService interface {
	EncodeURL(url models.EncodingRequest) *models.EncodedUrl
	DecodeURL(shortUrl models.DecodingRequest) (*models.DecodedUrl, *models.ErrorResponse)
}

type UrlServiceStruct struct {
	Store store.DataStore
}

func (urlEncoder *UrlServiceStruct) EncodeURL(url models.EncodingRequest) *models.EncodedUrl {
	// hash the original URL and add a constant prefix
	encodedString := constants.BaseURL + hash.GenerateHash(url.OriginalURL)
	// save the mapping of encoded and original URL in the store
	urlEncoder.Store.Save(store.Data{OriginalURL: url.OriginalURL, EncodedURL: encodedString})
	return &models.EncodedUrl{EncodedUrl: encodedString}
}

func (urlEncoder *UrlServiceStruct) DecodeURL(shortUrl models.DecodingRequest) (*models.DecodedUrl, *models.ErrorResponse) {
	// retrive the original URL from store which is mapped with the encoded URL sent in request
	originalUrl, err := urlEncoder.Store.Retrive(shortUrl.EncodedURL)
	if err != nil {
		return nil, &models.ErrorResponse{Message: err.Error()}
	}
	return &models.DecodedUrl{DecodedUrl: originalUrl}, nil
}
