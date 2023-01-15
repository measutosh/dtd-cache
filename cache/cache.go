package cache

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
  lock sync.RWMutex
  data map[string][]byte
}

func New() *Cache {
  return &Cache{
  	// lock: sync.RWMutex{},
  	data: make(map[string][]byte),
  }
}


func (c *Cache) Delete(key []byte) error {
  c.lock.RLock()
  defer c.lock.RUnlock()

  delete(c.data, string(key))

  return nil
}


func (c *Cache) Has(key []byte) bool {
  c.lock.RLock()
  defer c.lock.RLock() 

  _, ok := c.data[string(key)]
  return ok
}

func (c *Cache) Get(key []byte) ([]byte, error) {
  c.lock.RLock()
  defer c.lock.RUnlock()

  keyStr := string(key)

  val, ok := c.data[keyStr]
  if !ok {
    return nil, fmt.Errorf("key (%s) was not found", keyStr)
  }

  return val, nil
}

func (c *Cache) Set(key, value []byte, ttl time.Duration) error {
  c.lock.Lock()
  defer c.lock.Unlock()

  c.data[string(key)] = value

  return nil
}