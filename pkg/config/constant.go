package config

import "time"

const (
	HeartbeatEndpoint             = "/health"
	APITimeoutDuration            = 2 * time.Second
	ServerGracefulShutdownTimeout = 5 * time.Second

	CacheKey               = "mediumBlogPosts"
	CacheDefaultExpiration = 24 * time.Hour
	CacheCleanupInterval   = time.Hour

	LogTimestampFormat = "2006-01-02 15:04:05"
	LogFulltimestamp   = true
)
