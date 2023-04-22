package models

type EncodingRequest struct {
	OriginalURL string `json:"originalUrl"`
}

type DecodingRequest struct {
	EncodedURL string `json:"encodedUrl"`
}
