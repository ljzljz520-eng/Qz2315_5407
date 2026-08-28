package domain

import (
	"encoding/json"
	"time"
)

func EncodeRecord(r Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (Record, error) { var r Record; e := json.Unmarshal(b, &r); return r, e }
func Touch(r *Record, now time.Time) {
	if r.CreatedAt.IsZero() {
		r.CreatedAt = now
	}
	r.UpdatedAt = now
}
func CloneRecord(r Record) Record { return r }
func EventKinds() []string        { return []string{"submitted", "reviewed", "withdrawn", "notification"} }
