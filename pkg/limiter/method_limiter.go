package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"strings"
)

type MethodLimiter struct {
	*Limiter
}

func (l MethodLimiter) Key(c *gin.Context) string {
	uri := c.Request.RequestURI
	index := strings.Index(uri, "?")
	if index == -1 {
		return uri
	}

	return uri[:index]
}

func (l MethodLimiter) GetBucket(key string) (*ratelimit.Bucket, bool) {
	bucket, ok := l.LimiterBuckets[key]
	return bucket, ok
}

func (l MethodLimiter) AddBuckets(rules ...LimiterBucketRule) LimiterInterface {
	for _, rule := range rules {
		if _, ok := l.LimiterBuckets[rule.Key]; !ok {
			l.LimiterBuckets[rule.Key] = ratelimit.NewBucketWithQuantum(rule.FillInterval, rule.Capacity, rule.Quantum)
		}
	}
	return l
}

func NewMethodLimiter() LimiterInterface {
	return MethodLimiter{
		Limiter: &Limiter{LimiterBuckets: make(map[string]*ratelimit.Bucket)},
	}
}
