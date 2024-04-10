package main

import (
	"errors"
	"fmt"
	"sync"
)

//KeyValusStore represents a concurrent-safe key-value store
type KeyValueStore struct {
	data sync.Map
	sync.Mutex
}

//Creates a new value instance of  KeyValueStore
func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: sync.Map{}, //initialize with an empty sync Map
	}
}

/**
 * Puts a new key-value pair into the store
 * @param key
 * @param value
 * @return error
 */
func (kv *KeyValueStore) Put(key string, value string) error {
	kv.Lock()
	defer kv.Unlock()
	if _, ok := kv.data.Load(key); ok {
		return errors.New("key already exists")
	}
	kv.data.Store(key, value)
	return nil
}

/**
 * Fetch the value from store based on key
 * @param key
 * @return value
 * @return error
 */
func (kv *KeyValueStore) Get(key string) (string, error) {
	kv.Lock()
	defer kv.Unlock()
	value, ok := kv.data.Load((key))

	if !ok {
		return "", errors.New("key not found")
	}
	return value.(string), nil
}

/**
 * Remove key value associated with the given key from store
 * @param key
 * @return error
 */
func (kv *KeyValueStore) Delete(key string) error {
	kv.Lock()
	defer kv.Unlock()

	if _, ok := kv.data.Load(key); !ok {
		return errors.New("key not found")
	}
	kv.data.Delete(key)
	return nil
}

//Function to handle errors.
func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
