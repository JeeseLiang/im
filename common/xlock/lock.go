package xlock

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrLockFailed   = errors.New("获取锁失败")
	ErrUnlockFailed = errors.New("释放锁失败")
)

// Lock 分布式锁接口
type Lock interface {
	// Acquire 获取锁
	Acquire(ctx context.Context, key string, value string, ttl time.Duration) error
	// Release 释放锁
	Release(ctx context.Context, key string, value string) error
}

// RedisLock Redis实现的分布式锁
type RedisLock struct {
	client *redis.Client
}

// NewRedisLock 创建Redis分布式锁实例
func NewRedisLock(client *redis.Client) (*RedisLock, error) {
	return &RedisLock{
		client: client,
	}, nil
}

// acquireLockScript 获取锁的Lua脚本
const acquireLockScript = `
if redis.call('set', KEYS[1], ARGV[1], 'NX', 'PX', ARGV[2]) then
    return 1
else
    return 0
end
`

// releaseLockScript 释放锁的Lua脚本
const releaseLockScript = `
if redis.call('get', KEYS[1]) == ARGV[1] then
    return redis.call('del', KEYS[1])
else
    return 0
end
`

// Acquire 获取分布式锁
func (l *RedisLock) Acquire(ctx context.Context, key string, value string, ttl time.Duration) error {
	script := redis.NewScript(acquireLockScript)
	result, err := script.Run(ctx, l.client, []string{key}, value, ttl.Milliseconds()).Int()
	if err != nil {
		return err
	}
	if result == 0 {
		return ErrLockFailed
	}
	return nil
}

// Release 释放分布式锁
func (l *RedisLock) Release(ctx context.Context, key string, value string) error {
	script := redis.NewScript(releaseLockScript)
	result, err := script.Run(ctx, l.client, []string{key}, value).Int()
	if err != nil {
		return err
	}
	if result == 0 {
		return ErrUnlockFailed
	}
	return nil
}
