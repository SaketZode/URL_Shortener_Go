package service

import (
	"urlshortener/constants"
	"urlshortener/hashgenerator"
	"urlshortener/models"
	"urlshortener/store"
)

var hash hashgenerator.HashGenerator = &hashgenerator.MD5HashGenerator{}

type UrlService interface {
	EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse)
	DecodeURL(shortUrl models.DecodingRequest) (*models.DecodedUrl, *models.ErrorResponse)
}

type UrlServiceStruct struct {
	Store store.DataStore
}

func (urlEncoder *UrlServiceStruct) EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse) {
	encodedString := constants.BaseURL + hash.GenerateHash(url.OriginalURL)
	urlEncoder.Store.Save(store.Data{OriginalURL: url.OriginalURL, EncodedURL: encodedString})
	return &models.EncodedUrl{EncodedUrl: encodedString}, nil
}

func (urlEncoder *UrlServiceStruct) DecodeURL(shortUrl models.DecodingRequest) (*models.DecodedUrl, *models.ErrorResponse) {
	originalUrl, err := urlEncoder.Store.Retrive(shortUrl.EncodedURL)
	if err != nil {
		return nil, &models.ErrorResponse{Message: err.Error()}
	}
	return &models.DecodedUrl{DecodedUrl: originalUrl}, nil
}
