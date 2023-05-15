package cache

type Cache interface {
	// Set sets an item to the cache, replacing any existing item.
	Set(key string, value any)
	// Get gets an item from the cache.
	Get(key string) (any, bool)
}
