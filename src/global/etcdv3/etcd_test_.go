package etcdv3

import (
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func TestEtcdV3() {
	var wg sync.WaitGroup
	wg.Add(1)
	manager := NewManager()

	dis, _ := NewDiscovery(&Key{
		Name: "test",
		Addr: "127.0.0.1:8888",
	}, clientv3.Config{
		Endpoints:   []string{"192.168.0.113:2379"},
		DialTimeout: 5 * time.Second,
	}, manager)

	reg, _ := NewRegister(&Key{
		Key:  "discovery/testsvr/instance_id/aaabbbccc",
		Name: "testsvr",
		Addr: "127.0.0.1:8888",
	}, clientv3.Config{
		Endpoints:   []string{"192.168.0.113:2379"},
		DialTimeout: 5 * time.Second,
	})

	reg2, _ := NewRegister(&Key{
		Key:  "discovery/testsvr/instance_id/testqqqqq",
		Name: "testsvr",
		Addr: "127.0.0.1:9999",
	}, clientv3.Config{
		Endpoints:   []string{"192.168.0.113:2379"},
		DialTimeout: 5 * time.Second,
	})
	go reg.Run()
	time.Sleep(time.Second * 2)
	dis.pull()
	go dis.watch()
	time.Sleep(time.Second * 1)
	go reg2.Run()
	time.Sleep(time.Second * 1)
	manager.Dump()
	wg.Wait()
}
