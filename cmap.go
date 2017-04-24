package paxi

import "sync"

type cmap struct {
	data map[interface{}]interface{}
	sync.RWMutex
}

func NewCMap() *cmap {
	return &cmap{
		data: make(map[interface{}]interface{}),
	}
}

func (c *cmap) get(key interface{}) interface{} {
	c.RLock()
	defer c.RUnlock()
	return c.data[key]
}

func (c *cmap) set(key, value interface{}) {
	c.Lock()
	defer c.Unlock()
	c.data[key] = value
}

func (c *cmap) exist(key interface{}) bool {
	c.RLock()
	defer c.RUnlock()
	_, exist := c.data[key]
	return exist
}
