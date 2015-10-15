package memstore

import (
	"sync"

	"github.com/pressly/chainstore/lrumgr"
)

type memStore struct {
	sync.RWMutex
	data map[string][]byte
}

func New(capacity int64) *lrumgr.LruManager {
	memStore := &memStore{data: make(map[string][]byte, 1000)}
	store := lrumgr.New(capacity, memStore)
	return store
}

func (s *memStore) Open() (err error)  { return }
func (s *memStore) Close() (err error) { return }

func (s *memStore) Put(key string, val []byte) (err error) {
	s.Lock()
	defer s.Unlock()

	s.data[key] = val
	return nil
}

func (s *memStore) Get(key string) ([]byte, error) {
	s.RLock()
	defer s.RUnlock()

	val := s.data[key]
	return val, nil
}

func (s *memStore) Del(key string) (err error) {
	s.Lock()
	defer s.Unlock()

	delete(s.data, key)
	return
}
