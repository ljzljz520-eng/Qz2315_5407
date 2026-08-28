package store

import (
	"encoding/json"
	"go.etcd.io/bbolt"
	"independentjournal/domain"
	"os"
	"sync"
	"time"
)

var buckets = [][]byte{[]byte("records"), []byte("users"), []byte("events"), []byte("audits")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, &bbolt.Options{Timeout: time.Second})
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func put(tx *bbolt.Tx, b []byte, key string, v any) error {
	raw, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return tx.Bucket(b).Put([]byte(key), raw)
}
func get(tx *bbolt.Tx, b []byte, key string, v any) error {
	raw := tx.Bucket(b).Get([]byte(key))
	if raw == nil {
		return os.ErrNotExist
	}
	return json.Unmarshal(raw, v)
}
func (s *Store) SaveRecord(r domain.Record) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[0], r.ID, r) })
}
func (s *Store) LoadRecord(id string) (domain.Record, error) {
	var r domain.Record
	e := s.db.View(func(tx *bbolt.Tx) error { return get(tx, buckets[0], id, &r) })
	return r, e
}
func (s *Store) SaveUser(u domain.User) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[1], u.ID, u) })
}
func (s *Store) SaveEvent(v domain.Event) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[2], v.ID, v) })
}
func (s *Store) SaveAudit(v domain.Audit) error {
	return s.db.Update(func(tx *bbolt.Tx) error { return put(tx, buckets[3], v.ID, v) })
}
func (s *Store) ListEvents(rid string) ([]domain.Event, error) {
	out := []domain.Event{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[2]).ForEach(func(_, v []byte) error {
			var x domain.Event
			if v == nil {
				return nil
			}
			if err := json.Unmarshal(v, &x); err != nil {
				return err
			}
			if x.RecordID == rid {
				out = append(out, x)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) ListAudits(rid string) ([]domain.Audit, error) {
	out := []domain.Audit{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[3]).ForEach(func(_, v []byte) error {
			var x domain.Audit
			if v == nil {
				return nil
			}
			if err := json.Unmarshal(v, &x); err != nil {
				return err
			}
			if x.RecordID == rid {
				out = append(out, x)
			}
			return nil
		})
	})
	return out, e
}
