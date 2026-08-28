package journal

import (
	"fmt"
	"independentjournal/domain"
	"sort"
	"strings"
)

func RenderStatus(r domain.Record, events []domain.Event, audits []domain.Audit) string {
	sort.Slice(events, func(i, j int) bool { return events[i].CreatedAt.Before(events[j].CreatedAt) })
	sort.Slice(audits, func(i, j int) bool { return audits[i].CreatedAt.Before(audits[j].CreatedAt) })
	var b strings.Builder
	fmt.Fprintf(&b, "ID: %s\nTitle: %s\nBody: %s\nStatus: %s\nOwner: %s\n", r.ID, r.Title, r.Body, r.Status, r.OwnerID)
	for _, e := range events {
		fmt.Fprintf(&b, "Event: %s %s\n", e.Type, e.Payload)
	}
	for _, a := range audits {
		fmt.Fprintf(&b, "Audit: %s %s\n", a.Action, a.Detail)
	}
	return b.String()
}
func StatusLabel(r domain.Record) string {
	switch r.Status {
	case "withdrawn":
		return "Withdrawn"
	case "archived":
		return "Archived"
	default:
		return strings.Title(r.Status)
	}
}
func IsBlank(v string) bool { return strings.TrimSpace(v) == "" }
