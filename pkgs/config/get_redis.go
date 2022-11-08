// Butterfly - The web crawler base on Apache Solr for StarStart!
// Copyright(c) 2020 Star Inc. All Rights Reserved.
// The software licensed under Mozilla Public License Version 2.0

package config

import (
	"github.com/go-redis/redis/v8"
)

func GetRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     RedisAddress,
		Password: RedisPassword,
		DB:       RedisDatabase,
	})
}
