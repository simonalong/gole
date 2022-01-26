package etcd

import (
	"context"
	"errors"
	"fmt"
	"github.com/simonalong/gole/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdOrm struct {

	// 模块，也叫层次，默认为core
	Model string
	// 服务名
	ServiceName string
	// key的前缀，是 Model ServiceName的结合，为 /{Model}/{ServiceName}/
	KeyPre  string
	Client  *clientv3.Client
	Context context.Context
}

// ConnectDefault 默认链接
func ConnectDefault(model, serviceName string, endpoints []string, user, password string) (*EtcdOrm, error) {
	return ConnectConfig(model, serviceName, clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
		Username:    user,
		Password:    password,
	})
}

// ConnectConfig 创建实例
func ConnectConfig(model, serviceName string, etcdConfig clientv3.Config) (*EtcdOrm, error) {
	if "" == model {
		model = "core"
	}

	if "" == serviceName {
		return nil, errors.New("serviceName can't nil")
	}
	var clientV *clientv3.Client
	var err error
	if clientV, err = clientv3.New(etcdConfig); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &EtcdOrm{Model: model, ServiceName: serviceName, KeyPre: "/" + model + "/" + serviceName + "/", Client: clientV, Context: context.Background()}, nil
}

// Insert 新增
func (orm EtcdOrm) Insert(key string, value interface{}) error {
	_, err := orm.Client.Put(orm.Context, orm.KeyPre+key, util.ToJsonString(value))
	if err != nil {
		return err
	}
	return nil
}

// Delete 删除
func (orm EtcdOrm) Delete(key string, kvkvs ...string) error {
	_, err := orm.Client.Delete(orm.Context, orm.KeyPre+key)
	if err != nil {
		return err
	}
	return nil
}

// Update 修改
func (orm EtcdOrm) Update(key string, newValue interface{}) error {
	_, err := orm.Client.Put(orm.Context, orm.KeyPre+key, util.ToJsonString(newValue))
	if err != nil {
		return err
	}
	return nil
}

// One 查询：单行数据（单个实体数据）
func (orm EtcdOrm) One(key string) (interface{}, error) {
	rsp, err := orm.Client.Get(orm.Context, orm.KeyPre+key)
	if err != nil {
		return nil, err
	}
	var kvs = rsp.Kvs
	if nil != kvs {
		return kvs[0].Value, nil
	}
	return nil, nil
}

// 查询：多行数据（多个实体数据）
func (orm EtcdOrm) List() {
	// todo
}

// 查询：分页查询（多个实体数据）
func (orm EtcdOrm) Page() {
	// todo
}

// 查询：个数
func (orm EtcdOrm) Count() {
	// todo
}

// 查询：存在否
func (orm EtcdOrm) Exist() {
	// todo
}

// 查询：单个值
func (orm EtcdOrm) Value() {
	// todo
}

// 查询：多个值
func (orm EtcdOrm) Values() {
	// todo
}

// 查询：keys
func (orm EtcdOrm) Keys(key string) {
	// todo
	return
}

// 批处理：批量插入
func (orm EtcdOrm) BatchInsert() {
	// todo
}

// 批处理：批量更新
func (orm EtcdOrm) BatchUpdate() {
	// todo
}

// 事务
func (orm EtcdOrm) Tx() {
	//_, err := orm.Client.Txn(orm.Context)
	//if err != nil {
	//	return err
	//}
	//return nil
}

// 分布式锁：加锁
func (orm EtcdOrm) Lock() {
	// todo
}

// 分布式锁：解锁
func (orm EtcdOrm) UnLock() {
	// todo
}

// 监听
func (orm EtcdOrm) Watch() {
	// todo
}
