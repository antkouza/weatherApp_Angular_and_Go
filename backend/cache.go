package main

import (
	"fmt"
	"time"
)

func NewWeatherCache() *WeatherCache {
	return &WeatherCache{
		data: make(map[string]cacheItem),
	}
}

func (c *WeatherCache) Get(city string) (map[string]interface{}, bool) {
	c.mutex.RLock()
	item, found := c.data[city]
	c.mutex.RUnlock()

	if !found {
		return nil, false
	}

	if time.Now().After(item.expiresAt) {
		// Optimization: Remove entry now so it doesn't waste RAM
		c.mutex.Lock()
		delete(c.data, city)
		c.mutex.Unlock()

		fmt.Printf("Cache expired & deleted: %s\n", city)
		return nil, false
	}

	return item.data, true
}

func (c *WeatherCache) Set(city string, value map[string]interface{}, ttl time.Duration) {
	c.mutex.Lock()
	c.data[city] = cacheItem{
		data:      value,
		expiresAt: time.Now().Add(ttl),
	}
	c.mutex.Unlock()
}