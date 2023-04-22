package store

import "fmt"

type Data struct {
	OriginalURL string
	EncodedURL  string
}

type DataStore interface {
	Save(Data)
	Retrive(string) (string, error)
}

type Database struct {
	Data map[string]string
}

func NewDatabase() *Database {
	return &Database{Data: make(map[string]string)}
}

func (db *Database) Save(data Data) {
	db.Data[data.EncodedURL] = data.OriginalURL
}

func (db *Database) Retrive(encodedUrl string) (string, error) {
	/* for _, value := range db.Data {
		if value.EncodedURL == encodedUrl {
			return value.OriginalURL, nil
		}
	} */
	value, ok := db.Data[encodedUrl]
	if !ok {
		return "", fmt.Errorf("no mapping present for provided short URL")
	}
	return value, fmt.Errorf("no mapping present for provided short URL")
}
