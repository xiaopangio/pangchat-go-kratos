package cache

import "github.com/gorilla/websocket"

type ConnectionCache struct {
	*MemoryCache
}

func NewConnectionCache() *ConnectionCache {
	return &ConnectionCache{NewMemoryCache()}
}

func (c *ConnectionCache) GetConn(key string) (*websocket.Conn, bool) {
	v, b := c.MemoryCache.Get(key)
	if !b {
		return nil, false
	}
	return v.(*websocket.Conn), true
}
func (c *ConnectionCache) SetConn(key string, value *websocket.Conn) {
	c.MemoryCache.Set(key, value)
}
func (c *ConnectionCache) RemoveConn(key string) {
	c.MemoryCache.Remove(key)
}
