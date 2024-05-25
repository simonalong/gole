package test

import (
	"context"
	"fmt"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/extend/etcd"
	"github.com/simonalong/gole/time"
	"github.com/simonalong/gole/util"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func Test1(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	if config.GetValueBoolDefault("gole.etcd.enable", false) {
		err := config.GetValueObject("gole.etcd", &config.EtcdCfg)
		if err != nil {
			return
		}
	}

	etcdClient, _ := etcd.NewEtcdClient()

	ctx := context.Background()
	etcdClient.Put(ctx, "test", time.TimeToStringYmdHms(time.Now()))
	rsp, _ := etcdClient.Get(ctx, "test")
	etcdClient.Get(ctx, "test", func(pOp *clientv3.Op) {
		fmt.Println("信息")
		fmt.Println(util.ToJsonString(&pOp))
	})
	fmt.Println(rsp)
}

func TestRetry(t *testing.T) {
	config.LoadYamlFile("./application-retry.yaml")
	if config.GetValueBoolDefault("gole.etcd.enable", false) {
		err := config.GetValueObject("gole.etcd", &config.EtcdCfg)
		if err != nil {
			return
		}
	}

	_, err := etcd.NewEtcdClient()
	if err != nil {
		fmt.Println("")
	} else {
		fmt.Println("链接etcd 成功")
	}
}
