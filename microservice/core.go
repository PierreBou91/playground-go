package main

import (
	"errors"
	"sync"
)

var store = sync.Map{}

var ErrorNoSuchKey = errors.New("no such key")

func Put(key string, value string) error {
	store.Store(key, value)
	return nil
}

func Get(key string) (string, error) {
	value, ok := store.Load(key)
	if !ok {
		return "", ErrorNoSuchKey
	}
	return value.(string), nil
}

func Delete(key string) error {
	store.Delete(key)
	return nil
}
