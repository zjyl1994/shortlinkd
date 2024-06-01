package code

import (
	"sync"
	"time"
)

var (
	data = make(map[string]CodeItem)
	lock sync.RWMutex
)

type CodeItem struct {
	Code    string
	URL     string
	Enabled bool
	Expired *time.Time
}
