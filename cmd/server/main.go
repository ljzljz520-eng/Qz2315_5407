package main

import (
	"flag"
	"independentjournal/api"
	"independentjournal/service"
	"independentjournal/store"
	"log"
	"net/http"
)

func main() {
	path := flag.String("db", "journal.db", "database path")
	addr := flag.String("addr", ":8080", "listen address")
	flag.Parse()
	st, e := store.Open(*path)
	if e != nil {
		log.Fatal(e)
	}
	defer st.Close()
	svc := service.New(st)
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, api.NewHandler(svc)))
}
