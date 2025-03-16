package services

import "sync"

var (
	index   = make(map[string]map[string]int)
	visited = make(map[string]bool)
	mu      sync.Mutex
)
