package service

import (
	"fmt"
	"independentjournal/domain"
	"independentjournal/journal"
	"independentjournal/notify"
	"independentjournal/store"
	"time"
)

type Service struct {
	Store *store.Store
	Clock func() time.Time
	Seq   int
}

func New(s *store.Store) *Service          { return &Service{Store: s, Clock: time.Now} }
func (s *Service) id(prefix string) string { s.Seq++; return fmt.Sprintf("%s-%03d", prefix, s.Seq) }
func (s *Service) RegisterEntry(r domain.Record, u domain.User) (domain.Record, error) {
	r.Title = domain.NormalizeTitle(r.Title)
	r.Body = domain.NormalizeBody(r.Body)
	if err := r.Validate(); err != nil {
		return r, err
	}
	if !domain.ValidRole(u.Role) {
		return r, fmt.Errorf("invalid role")
	}
	now := s.Clock()
	r.Status = "submitted"
	r.UpdatedAt = now
	if err := s.Store.SaveRecord(r); err != nil {
		return r, err
	}
	if err := s.Store.SaveUser(u); err != nil {
		return r, err
	}
	_ = s.Store.SaveEvent(domain.NewEvent(s.id("evt"), r.ID, "submitted", r.Summary(), now))
	_ = s.Store.SaveAudit(domain.NewAudit(s.id("aud"), r.ID, u.ID, "register", "submitted", now))
	return r, nil
}
func (s *Service) ReviewEntry(id, actor, decision string) (domain.Record, error) {
	r, e := s.Store.LoadRecord(id)
	if e != nil {
		return r, e
	}
	if !domain.CanReview(r, "moderator") {
		return r, fmt.Errorf("review denied")
	}
	now := s.Clock()
	if decision != "approved" && decision != "rejected" && decision != "archived" {
		return r, fmt.Errorf("unsupported decision")
	}
	if e = r.Transition(decision, now); e != nil {
		return r, e
	}
	if e = s.Store.SaveRecord(r); e != nil {
		return r, e
	}
	_ = s.Store.SaveAudit(domain.NewAudit(s.id("aud"), id, actor, "review", decision, now))
	return r, nil
}
func (s *Service) WithdrawEntry(id, actor string) (domain.Record, error) {
	r, e := s.Store.LoadRecord(id)
	if e != nil {
		return r, e
	}
	if !domain.CanWithdraw(r, actor) {
		return r, fmt.Errorf("withdraw denied")
	}
	now := s.Clock()
	if e = r.Transition("withdrawn", now); e != nil {
		return r, e
	}
	if e = s.Store.SaveRecord(r); e != nil {
		return r, e
	}
	_ = notify.SendWithdrawalNotice(s.Store, r, actor, now, s.id)
	return r, nil
}
func (s *Service) QueryEntry(id string) (string, error) {
	r, e := s.Store.LoadRecord(id)
	if e != nil {
		return "", e
	}
	ev, _ := s.Store.ListEvents(id)
	au, _ := s.Store.ListAudits(id)
	return journal.RenderStatus(r, ev, au), nil
}
