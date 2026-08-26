package service

import "independentjournal/domain"

func NextReviewState(decision string) string {
	if decision == "approve" || decision == "approved" {
		return "approved"
	}
	if decision == "reject" || decision == "rejected" {
		return "rejected"
	}
	return "archived"
}
func CanQuery(r domain.Record, role string) bool {
	if role == "moderator" || role == "reader" {
		return true
	}
	return r.OwnerID == role
}
func StatusComplete(r domain.Record) bool { return r.IsComplete() }
func NormalizeDecision(v string) string {
	if v == "approve" {
		return "approved"
	}
	if v == "reject" {
		return "rejected"
	}
	return v
}
