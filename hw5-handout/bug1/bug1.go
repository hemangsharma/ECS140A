package bug1

import "sync"

var he sync.Mutex

// Counter stores a count.
type Counter struct {
	n int64
}

// Inc increments the count in the Counter.
func (c *Counter) Inc() {
	he.Lock()
	c.n++
	he.Unlock()
}
