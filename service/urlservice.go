package service

import (
	"crypto/md5"
	"encoding/hex"
	"urlshortener/constants"
	"urlshortener/models"
	"urlshortener/store"
)

type UrlService interface {
	EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse)
	DecodeURL(shortUrl models.DecodingRequest) (*models.DecodedUrl, *models.ErrorResponse)
}

type UrlServiceStruct struct {
	Store store.DataStore
}

func (urlEncoder *UrlServiceStruct) EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse) {
	hash := md5.Sum([]byte(url.OriginalURL))
	encodedString := constants.BaseURL + hex.EncodeToString(hash[:])
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
