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

/*
Struct to store the data for encoded and original URL.
The data will be stored as a map where encoded(short) URL
will be the key and original URL will be the value associated
with the key.
*/
type Database struct {
	Data map[string]string
}

// allocate memory to the map declared in Database struct
func NewDatabase() *Database {
	return &Database{Data: make(map[string]string)}
}

// save the data in map
func (db *Database) Save(data Data) {
	db.Data[data.EncodedURL] = data.OriginalURL
}

// retrive the data from map
func (db *Database) Retrive(encodedUrl string) (string, error) {
	value, ok := db.Data[encodedUrl]
	if !ok {
		return "", fmt.Errorf("no mapping present for provided short URL")
	}
	return value, nil
}
