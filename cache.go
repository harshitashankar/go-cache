package cache

import (
	"sync"
	"errors"
)

type Key string
type Value interface{}

type Cache struct{
	data map[Key]Value
	mutex sync.Mutex
}

const threaded = true

func New() *Cache {
	c := &Cache{
		data: make(map[Key]Value),
	}
	return c
}

func(c *Cache) Get(key Key) (Value, bool){
	if threaded{
		c.mutex.Lock()
		defer c.mutex.Unlock()
	}

	value, exists := c.data[key]
	if !exists{
		return nil, false
	}
	return value, true
}

func(c *Cache) Set(key Key, value interface{}) {
	if threaded{
		c.mutex.Lock()
		defer c.mutex.Unlock()
	}
	c.data[key] = value
}

func(c *Cache) Remove(k Key) (err error){
	if threaded{
		c.mutex.Lock()
		defer c.mutex.Unlock()
	}

	err = nil
	defer func() error {
		if r := recover(); r != nil {
			err = errors.New("Remove: panicked!")
		}
		return err
	}()
	delete(c.data, k)	
	//panic("test1")
	return err
}