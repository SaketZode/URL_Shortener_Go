package models

// Response to be sent after encoding the original URL
type EncodedUrl struct {
	EncodedUrl string `json:"encodedUrl"`
}

// Response to be sent for decoding the encoded URL
type DecodedUrl struct {
	DecodedUrl string `json:"decodedUrl"`
}
