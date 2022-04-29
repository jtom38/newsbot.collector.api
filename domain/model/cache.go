package model

import (
	"time"
)

type CacheItem struct {
	Key   string
	Value string

	// Group defines what it should be a reference to.
	// youtube, reddit, ect
	Group   string
	Expires time.Time
	IsTainted bool
}