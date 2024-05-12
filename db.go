package kvdb

import (
	"errors"
	"sync"
)

type KVDB interface {
	Get(key string) (any, error)
	Put(key string, value any) error
	Delete(key string) error
	GetAllKeys() ([]string, error)
	Exists(key string) (bool, error)
}

type db struct {
	m     map[string]any
	mutex sync.RWMutex
}

func New() KVDB {
	return &db{
		m:     make(map[string]any),
		mutex: sync.RWMutex{},
	}
}

func (k *db) Get(key string) (any, error) {
	k.mutex.RLock()
	if val, ok := k.m[key]; ok {
		return val, nil
	}
	k.mutex.RUnlock()
	return nil, errors.New("key not found")
}

func (k *db) Put(key string, value any) error {
	k.mutex.Lock()
	k.m[key] = value
	k.mutex.Unlock()
	return nil
}

func (k *db) Delete(key string) error {
	k.mutex.Lock()
	delete(k.m, key)
	k.mutex.Unlock()
	return nil
}

func (k *db) GetAllKeys() ([]string, error) {
	var i int
	k.mutex.Lock()
	keys := make([]string, len(k.m))
	for key := range k.m {
		keys[i] = key
		i++
	}
	k.mutex.Unlock()
	return keys, nil
}

func (k *db) Exists(key string) (bool, error) {
	if _, ok := k.m[key]; ok {
		return true, nil
	}
	return false, nil
}
