package notify

import (
	"independentjournal/domain"
	"independentjournal/store"
	"time"
)

func SendWithdrawalNotice(s *store.Store, r domain.Record, actor string, now time.Time, nextID func(string) string) error {
	e := domain.NewEvent(nextID("evt"), r.ID, "withdrawal_notice", "owner="+r.OwnerID, now)
	if err := s.SaveEvent(e); err != nil {
		return err
	}
	return s.SaveAudit(domain.NewAudit(nextID("aud"), r.ID, actor, "notify", "withdrawal notice sent", now))
}
func ChannelForRole(role string) string {
	if role == "moderator" {
		return "moderation"
	}
	return "journal"
}
func DeliveryMessage(r domain.Record) string { return "Withdrawal processed for " + r.Title }
