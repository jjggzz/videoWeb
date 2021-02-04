package dao

import "github.com/gomodule/redigo/redis"

// 设置带过期时间的缓存(单位:秒)
func (dao *Dao) SetexRedisCache(timeout int64, key string, value string) error {
	conn := dao.redis.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("setex", key, timeout, value)
	return err
}

// 获取缓存
func (dao *Dao) GetRedisCache(key string) (string, error) {
	conn := dao.redis.Get()
	defer func() {
		_ = conn.Close()
	}()
	str, err := redis.String(conn.Do("get", key))
	if err != nil && err != redis.ErrNil {
		return "", err
	}
	return str, nil
}

// 删除缓存
func (dao *Dao) DelRedisCache(key string) error {
	conn := dao.redis.Get()
	defer func() {
		_ = conn.Close()
	}()
	_, err := conn.Do("del", key)
	return err
}
