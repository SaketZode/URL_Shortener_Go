package models

// Request for encoding the URL
type EncodingRequest struct {
	OriginalURL string `json:"originalUrl"`
}

// Request for decoding the URL
type DecodingRequest struct {
	EncodedURL string `json:"encodedUrl"`
}
