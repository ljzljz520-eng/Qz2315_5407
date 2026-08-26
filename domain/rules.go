package domain

import "strings"

func NormalizeTitle(v string) string { return strings.TrimSpace(v) }
func NormalizeBody(v string) string  { return strings.TrimSpace(v) }
func ValidRole(v string) bool {
	switch v {
	case "developer", "moderator", "reader":
		return true
	default:
		return false
	}
}
func ValidStatus(v string) bool {
	switch v {
	case "draft", "submitted", "approved", "rejected", "withdrawn", "archived":
		return true
	default:
		return false
	}
}
func CanReview(r Record, role string) bool {
	return role == "moderator" && (r.Status == "submitted" || r.Status == "approved" || r.Status == "rejected")
}
func CanWithdraw(r Record, actor string) bool { return actor == r.OwnerID || actor == "moderator" }
