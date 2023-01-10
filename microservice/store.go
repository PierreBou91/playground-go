package main

import (
	"errors"
	"log"
	"net/http"
)

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func Put(key string, value string) error {
	store[key] = value

	return nil
}

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Delete(key string) error {
	delete(store, key)
	return nil
}

func HelloGoHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello net/http!\n"))
}

func main() {
	http.HandleFunc("/", HelloGoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
