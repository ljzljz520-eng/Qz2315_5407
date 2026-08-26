package api

import (
	"encoding/json"
	"independentjournal/domain"
	"independentjournal/service"
	"net/http"
	"strings"
)

type Handler struct{ S *service.Service }

func NewHandler(s *service.Service) *Handler { return &Handler{S: s} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 || parts[0] != "records" {
		http.NotFound(w, r)
		return
	}
	id := parts[1]
	switch {
	case r.Method == http.MethodGet && len(parts) == 2:
		h.get(w, id)
	case r.Method == http.MethodPost && len(parts) == 3 && parts[2] == "withdraw":
		h.withdraw(w, r, id)
	default:
		http.NotFound(w, r)
	}
}
func (h *Handler) get(w http.ResponseWriter, id string) {
	v, e := h.S.QueryEntry(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	_, _ = w.Write([]byte(v))
}
func (h *Handler) withdraw(w http.ResponseWriter, r *http.Request, id string) {
	var in struct {
		Actor string `json:"actor"`
	}
	_ = json.NewDecoder(r.Body).Decode(&in)
	v, e := h.S.WithdrawEntry(id, in.Actor)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	_ = json.NewEncoder(w).Encode(v)
}
func DecodeRecord(r *http.Request) (domain.Record, error) {
	var x domain.Record
	e := json.NewDecoder(r.Body).Decode(&x)
	return x, e
}
