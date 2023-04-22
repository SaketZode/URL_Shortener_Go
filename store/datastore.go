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
	Data []Data
}

func (db *Database) Save(data Data) {
	db.Data = append(db.Data, data)
}

func (db *Database) Retrive(encodedUrl string) (string, error) {
	for _, value := range db.Data {
		if value.EncodedURL == encodedUrl {
			return value.OriginalURL, nil
		}
	}
	return "", fmt.Errorf("no mapping present for provided short URL")
}
