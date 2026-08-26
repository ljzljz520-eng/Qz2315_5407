package journal

import (
	"independentjournal/domain"
	"sort"
)

func LatestEvent(events []domain.Event) domain.Event {
	if len(events) == 0 {
		return domain.Event{}
	}
	sort.Slice(events, func(i, j int) bool { return events[i].CreatedAt.After(events[j].CreatedAt) })
	return events[0]
}
func HasType(events []domain.Event, t string) bool {
	for _, e := range events {
		if e.Type == t {
			return true
		}
	}
	return false
}
func AuditCount(a []domain.Audit) int { return len(a) }
func Timeline(events []domain.Event) []string {
	out := make([]string, 0, len(events))
	for _, e := range events {
		out = append(out, e.Type)
	}
	return out
}
