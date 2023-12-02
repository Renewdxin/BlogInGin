package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// LimiterIface is an interface for rate limiting operations.
type LimiterIface interface {
	// Key generates a unique key for rate limiting based on the given Gin context.
	Key(c *gin.Context) string

	// GetBucket retrieves the rate limiting bucket for the specified key.
	// It returns the bucket and a boolean indicating whether the bucket was found.
	GetBucket(key string) (*ratelimit.Bucket, bool)

	// AddBuckets adds rate limiting buckets based on the provided rules.
	// It returns the modified LimiterIface with added buckets.
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

type Limiter struct {
	limiterBuckets map[string]*ratelimit.Bucket
}

type LimiterBucketRule struct {
	Key          string
	FillInterval time.Duration
	Capacity     int64
	Quantum      int64 //放置的令牌数量
}
