package service

import (
	"crypto/md5"
	"encoding/hex"
	"urlshortener/constants"
	"urlshortener/models"
	"urlshortener/store"
)

type UrlEncodeService interface {
	EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse)
}

type UrlEncoder struct {
	Store store.DataStore
}

func (urlEncoder *UrlEncoder) EncodeURL(url models.EncodingRequest) (*models.EncodedUrl, *models.ErrorResponse) {
	hash := md5.Sum([]byte(url.OriginalURL))
	encodedString := constants.BaseURL + hex.EncodeToString(hash[:])
	urlEncoder.Store.Save(store.Data{OriginalURL: url.OriginalURL, EncodedURL: encodedString})
	return &models.EncodedUrl{EncodedUrl: encodedString}, nil
}
