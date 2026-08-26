package domain

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var ErrInvalidRecord = errors.New("invalid record")
var ErrInvalidTransition = errors.New("invalid transition")

type Record struct {
	ID, Title, Body, Status, OwnerID string
	CreatedAt, UpdatedAt             time.Time
}
type User struct {
	ID, Name, Role string
	CreatedAt      time.Time
}
type Event struct {
	ID, RecordID, Type, Payload string
	CreatedAt                   time.Time
}
type Audit struct {
	ID, RecordID, ActorID, Action, Detail string
	CreatedAt                             time.Time
}

func NewRecord(id, title, body, owner string, now time.Time) Record {
	return Record{ID: id, Title: title, Body: body, Status: "draft", OwnerID: owner, CreatedAt: now, UpdatedAt: now}
}
func NewUser(id, name, role string, now time.Time) User {
	return User{ID: id, Name: name, Role: role, CreatedAt: now}
}
func (r Record) Validate() error {
	if strings.TrimSpace(r.ID) == "" || strings.TrimSpace(r.Title) == "" || strings.TrimSpace(r.Body) == "" || strings.TrimSpace(r.OwnerID) == "" {
		return ErrInvalidRecord
	}
	return nil
}
func (r Record) IsComplete() bool {
	return r.Title != "" && r.Body != "" && r.OwnerID != "" && r.Status != ""
}
func (r *Record) Transition(next string, now time.Time) error {
	allowed := map[string]map[string]bool{"draft": {"submitted": true}, "submitted": {"approved": true, "rejected": true, "withdrawn": true}, "approved": {"archived": true, "withdrawn": true}, "rejected": {"archived": true}, "withdrawn": {"archived": true}}
	if !allowed[r.Status][next] {
		return fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, r.Status, next)
	}
	r.Status = next
	r.UpdatedAt = now
	return nil
}
func (r Record) Summary() string {
	return fmt.Sprintf("%s|%s|%s|%s", r.ID, r.Title, r.Status, r.OwnerID)
}
func NewEvent(id, rid, typ, payload string, now time.Time) Event {
	return Event{ID: id, RecordID: rid, Type: typ, Payload: payload, CreatedAt: now}
}
func NewAudit(id, rid, actor, action, detail string, now time.Time) Audit {
	return Audit{ID: id, RecordID: rid, ActorID: actor, Action: action, Detail: detail, CreatedAt: now}
}
