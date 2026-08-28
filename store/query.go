package store

import "independentjournal/domain"

func (s *Store) SaveAll(r domain.Record, u domain.User, e domain.Event, a domain.Audit) error {
	if err := s.SaveRecord(r); err != nil {
		return err
	}
	if err := s.SaveUser(u); err != nil {
		return err
	}
	if err := s.SaveEvent(e); err != nil {
		return err
	}
	return s.SaveAudit(a)
}
func (s *Store) Exists(id string) bool     { _, e := s.LoadRecord(id); return e == nil }
func (s *Store) CountEvents(id string) int { v, _ := s.ListEvents(id); return len(v) }
func (s *Store) CountAudits(id string) int { v, _ := s.ListAudits(id); return len(v) }
