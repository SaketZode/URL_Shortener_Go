package service

import "crypto/md5"

type UrlEncodeService interface {
	EncodeURL(url string) (string, error)
}

type UrlEncoder struct {
}

func (urlEncoder *UrlEncoder) EncodeURL(url string) (string, error) {
	md5.Sum([]byte(url))
	// store.Save()
	return "", nil
}
