package etcdv3

import (
	"math/rand"
	"strings"
	"sync"

	"github.com/xi123/libgo/logs"
)

// <summary>
// Key
// <summary>
type Key struct {
	Key  string
	Name string
	Addr string
}

// <summary>
// Manager
// <summary>
type Manager struct {
	sync.RWMutex
	// <name,<id,Key>>
	m map[string]map[string]*Key
}

func NewManager() (m *Manager) {
	return &Manager{
		m: map[string]map[string]*Key{},
	}
}

func (s *Manager) Add(key *Key) {
	if key == nil {
		return
	}
	s.Lock()
	defer s.Unlock()
	if _, exist := s.m[key.Name]; !exist {
		s.m[key.Name] = map[string]*Key{}
	}
	s.m[key.Name][key.Key] = key
}

func (s *Manager) Del(id string) {
	logs.Debugf("%v", id)
	sli := strings.Split(id, "/")
	name := sli[len(sli)-2]
	s.Lock()
	defer s.Unlock()
	if _, exist := s.m[name]; exist {
		delete(s.m[name], id)
	}
}

func (s *Manager) Pick(name string) *Key {
	s.RLock()
	defer s.RUnlock()
	if m, exist := s.m[name]; !exist {
		return nil
	} else {
		// 纯随机取节点
		idx := rand.Intn(len(m))
		for _, v := range m {
			if idx == 0 {
				logs.Infof("%v %+v", name, v)
				return v
			}
			idx--
		}
	}
	return nil
}

func (s *Manager) Dump() {
	for key, val := range s.m {
		for k, v := range val {
			logs.Debugf("Name:%v Id:%v Key:%+v", key, k, v)
		}
	}
}
