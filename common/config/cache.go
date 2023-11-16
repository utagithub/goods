/*
 * @Author: 尤_Ta
 * @Date:  12:06
 * @Last Modified by: 尤_Ta
 * @Last Modified time:  12:06
 */

package config

import (
	"goods/common/storage"
	"goods/common/storage/cache"
)

type Cache struct {
	Redis *RedisConnectOptions
	//Memory interface{}//原
	//改造后
	Memory *CacheMemory
}

type CacheMemory struct {
	PoolSize uint
}

// CacheConfig cache配置
var CacheConfig = new(Cache)

// Setup 构造cache 顺序 redis > 其他 > memory
func (e Cache) Setup() (storage.AdapterCache, error) {
	if e.Redis != nil {
		options, err := e.Redis.GetRedisOptions()
		if err != nil {
			return nil, err
		}
		r, err := cache.NewRedis(GetRedisClient(), options)
		if err != nil {
			return nil, err
		}
		if _redis == nil {
			_redis = r.GetClient()
		}
		return r, nil
	}
	return cache.NewMemory(), nil
}
