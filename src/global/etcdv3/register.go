package etcdv3

import (
	"context"
	"encoding/json"
	"time"

	_ "github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"
	"github.com/xi123/libgo/logs"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/pkg/errors"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
)

const (
	_ttl = 10
)

// <summary>
// Register
// <summary>
type Register struct {
	cli       *clientv3.Client
	leaseId   clientv3.LeaseID
	lease     clientv3.Lease
	key       *Key
	closeChan chan error
}

func NewRegister(key *Key, conf clientv3.Config) (reg *Register, err error) {
	s := &Register{}
	s.closeChan = make(chan error)
	s.key = key
	s.cli, err = clientv3.New(conf)
	return s, err
}

func (s *Register) Run() {
	dur := time.Duration(time.Second)
	timer := time.NewTicker(dur)
	s.register()
	for {
		select {
		case <-timer.C:
			s.keepAlive()
		case <-s.closeChan:
			goto EXIT
		}
	}
EXIT:
	logs.Infof("exit...")
}

func (s *Register) Stop() {
	s.revoke()
	close(s.closeChan)
}

func (s *Register) register() (err error) {
	s.leaseId = 0
	kv := clientv3.NewKV(s.cli)
	s.lease = clientv3.NewLease(s.cli)
	leaseResp, err := s.lease.Grant(context.TODO(), _ttl)
	if err != nil {
		err = errors.New(logs.SprintErrorf(err.Error()))
		return
	}
	data, _ := json.Marshal(s.key)
	_, err = kv.Put(context.TODO(), s.key.Key, string(data), clientv3.WithLease(leaseResp.ID))
	if err != nil {
		err = errors.Wrap(err, logs.SprintErrorf("clientv3.KV.Put %v-%+v", s.key.Name, string(data)))
		return
	}
	s.leaseId = leaseResp.ID
	return
}

func (s *Register) keepAlive() (err error) {
	_, err = s.lease.KeepAliveOnce(context.TODO(), s.leaseId)
	if err != nil {
		if err == rpctypes.ErrLeaseNotFound {
			s.register()
			err = nil
		}
		err = errors.New(logs.SprintErrorf(err.Error()))
	}
	// logs.Infof("%+v", s.leaseId)
	return err
}

func (s *Register) revoke() (err error) {
	_, err = s.cli.Revoke(context.TODO(), s.leaseId)
	if err != nil {
		err = errors.New(logs.SprintErrorf(err.Error()))
		return
	}
	logs.Infof("%+v", s.leaseId)
	return
}
