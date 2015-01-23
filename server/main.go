package main

import (
	"libs/userdb"
	"log"
	"net/http"
)

const BindAddr string = ":8080"

// check against api keys
// write ip to database (api, ip, old ip)
func main() {

	if !userdb.Exists() {
		userdb.Create()
		log.Printf("Database created")
	}

	router := NewRouter()
	log.Printf("Listening on %s", BindAddr)
	log.Fatal(http.ListenAndServe(BindAddr, router))
}

// root iterate apis
// set up iptables, remove old ip, set nil in db
