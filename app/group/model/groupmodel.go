package model

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"im_message/common/biz"
	"im_message/common/xlock"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GroupModel = (*customGroupModel)(nil)

const (
	lockGroupPrefix = "lock:group:"
)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		TransInsertSystemGroup(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error)
		CreateGroup(ctx context.Context, session sqlx.Session, userId int64, groupName string) (string, error)
	}

	customGroupModel struct {
		*defaultGroupModel
		lock *xlock.RedisLock
	}
)

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn, c cache.CacheConf, lock *xlock.RedisLock) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn, c),
		lock:              lock,
	}
}

// 包装 Insert 方法
func (m *customGroupModel) Insert(ctx context.Context, data *Group) (sql.Result, error) {
	lockKey := fmt.Sprintf("%s%v", lockGroupPrefix, data.Id)

	// 获取分布式锁
	err := m.lock.Acquire(ctx, lockKey, data.Id, 10*time.Second)
	if err != nil {
		return nil, err
	}
	defer m.lock.Release(ctx, lockKey, data.Id)

	return m.defaultGroupModel.Insert(ctx, data)
}

// 包装 TransInsert 方法
func (m *customGroupModel) TransInsert(ctx context.Context, session sqlx.Session, data *Group) (sql.Result, error) {
	lockKey := fmt.Sprintf("%s%v", lockGroupPrefix, data.Id)

	// 获取分布式锁
	err := m.lock.Acquire(ctx, lockKey, data.Id, 10*time.Second)
	if err != nil {
		return nil, err
	}
	defer m.lock.Release(ctx, lockKey, data.Id)

	return m.defaultGroupModel.TransInsert(ctx, session, data)
}

// 包装 Update 方法
func (m *customGroupModel) Update(ctx context.Context, data *Group) error {
	lockKey := fmt.Sprintf("%s%v", lockGroupPrefix, data.Id)

	// 获取分布式锁
	err := m.lock.Acquire(ctx, lockKey, data.Id, 10*time.Second)
	if err != nil {
		return err
	}
	defer m.lock.Release(ctx, lockKey, data.Id)

	return m.defaultGroupModel.Update(ctx, data)
}

// 包装 TransUpdate 方法
func (m *customGroupModel) TransUpdate(ctx context.Context, session sqlx.Session, data *Group) error {
	lockKey := fmt.Sprintf("%s%v", lockGroupPrefix, data.Id)

	// 获取分布式锁
	err := m.lock.Acquire(ctx, lockKey, data.Id, 10*time.Second)
	if err != nil {
		return err
	}
	defer m.lock.Release(ctx, lockKey, data.Id)

	return m.defaultGroupModel.TransUpdate(ctx, session, data)
}

// 包装 Delete 方法
func (m *customGroupModel) Delete(ctx context.Context, id string) error {
	lockKey := fmt.Sprintf("%s%v", lockGroupPrefix, id)

	// 获取分布式锁
	err := m.lock.Acquire(ctx, lockKey, id, 10*time.Second)
	if err != nil {
		return err
	}
	defer m.lock.Release(ctx, lockKey, id)

	return m.defaultGroupModel.Delete(ctx, id)
}

func (m *defaultGroupModel) CreateGroup(ctx context.Context, session sqlx.Session, userId int64, groupName string) (string, error) {
	// 创建群聊群组
	group_id := biz.GetUuid()
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	_, err := session.ExecCtx(ctx, query, group_id, groupName, GroupTypeMulti, GroupStatusYes, nil)
	if err != nil {
		return "", err
	}
	return group_id, nil
}

// 添加系统用户 组
func (m *defaultGroupModel) TransInsertSystemGroup(ctx context.Context, session sqlx.Session, userId int64) (sql.Result, error) {
	// 创建 与 微信团队 的群
	groupId1 := biz.GetGroupId(1, userId)
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err := session.ExecCtx(ctx, query, groupId1, "微信团队", GroupTypeSingle, GroupStatusYes, nil)
	if err != nil {
		return nil, err
	}
	// 创建 与 文件传输助手 的群
	groupId2 := biz.GetGroupId(2, userId)
	query = fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, groupRowsExpectAutoSet)
	ret, err = session.ExecCtx(ctx, query, groupId2, "文件传输助手", GroupTypeSingle, GroupStatusYes, nil)
	if err != nil {
		return nil, err
	}
	return ret, err
}
