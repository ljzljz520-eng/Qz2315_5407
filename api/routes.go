package api

import "net/http"

func Routes(h http.Handler) *http.ServeMux              { m := http.NewServeMux(); m.Handle("/", h); return m }
func Health(w http.ResponseWriter, r *http.Request)     { w.WriteHeader(http.StatusNoContent) }
func MethodAllowed(r *http.Request, method string) bool { return r.Method == method }
