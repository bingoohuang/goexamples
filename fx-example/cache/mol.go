package cache

import (
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

const molCacheKey = "mol"

type MeaningOfLifeCache interface {
	LoadOrStore(originFunc func() (string, error)) (string, error)
}

type meaningOfLifeRedisCache struct {
	client *redis.Client
	logger *zap.Logger
}

func NewMeaningOfLifeCacheRedis(client *redis.Client, logger *zap.Logger) MeaningOfLifeCache {
	return &meaningOfLifeRedisCache{
		client: client,
		logger: logger,
	}
}

func (c *meaningOfLifeRedisCache) LoadOrStore(originFunc func() (string, error)) (res string, err error) {
	if res, err = c.client.Get(molCacheKey).Result(); err == nil {
		return
	}

	if err != redis.Nil {
		c.logger.Error(err.Error())
		return
	}

	if res, err = originFunc(); err != nil {
		c.logger.Error(err.Error())
		return
	}

	if err = c.client.Set(molCacheKey, res, time.Minute).Err(); err != nil {
		c.logger.Error(err.Error())
		return
	}

	return
}
