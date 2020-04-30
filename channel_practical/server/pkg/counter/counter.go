package counter

import "sync"

// Client is counter with atomic increment
type Client struct {
	count int64
	mu    sync.RWMutex
}

// New returns new Counter which starts from zero
func New() *Client {
	return &Client{
		count: 0,
		mu:    sync.RWMutex{},
	}
}

// Increment atomically increment counter
func (c *Client) Increment() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	return c.count
}

// Current returns current count
func (c *Client) Current() int64 {
	// 読み取りロック
	// 読み取りロック同士はロック待ちにならない
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}
