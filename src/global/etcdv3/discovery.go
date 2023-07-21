package etcdv3

import (
	"context"
	"encoding/json"

	_ "github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/xi123/libgo/logs"
	"github.com/pkg/errors"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// <summary>
// Discovery
// <summary>
type Discovery struct {
	cli *clientv3.Client
	key *Key
	m   *Manager
}

func NewDiscovery(key *Key, conf clientv3.Config, m *Manager) (dis *Discovery, err error) {
	d := &Discovery{}
	d.key = key
	if m == nil {
		return nil, errors.New(logs.SprintErrorf(""))
	}
	d.m = m
	d.cli, err = clientv3.New(conf)
	return d, err
}

func (s *Discovery) pull() {
	kv := clientv3.NewKV(s.cli)
	resp, err := kv.Get(context.TODO(), "discovery/", clientv3.WithPrefix())
	if err != nil {
		logs.Fatalf(err.Error())
		return
	}
	for _, v := range resp.Kvs {
		key := &Key{}
		err = json.Unmarshal(v.Value, key)
		if err != nil {
			logs.Fatalf(err.Error())
			continue
		}
		s.m.Add(key)
		logs.Infof("%+v", key)
	}
}

func (s *Discovery) watch() {
	watcher := clientv3.NewWatcher(s.cli)
	watchChan := watcher.Watch(context.TODO(), "discovery", clientv3.WithPrefix())
	for {
		select {
		case resp := <-watchChan:
			s.watchEvent(resp.Events)
		}
	}
}

func (s *Discovery) watchEvent(evs []*clientv3.Event) {
	for _, ev := range evs {
		switch ev.Type {
		case clientv3.EventTypePut:
			key := &Key{}
			err := json.Unmarshal(ev.Kv.Value, key)
			if err != nil {
				logs.Fatalf(err.Error())
				continue
			}
			s.m.Add(key)
			logs.Infof("new %v", string(ev.Kv.Value))
		case clientv3.EventTypeDelete:
			s.m.Del(string(ev.Kv.Key))
			logs.Infof("del %v %v", string(ev.Kv.Key), string(ev.Kv.Value))
		}
	}
}
