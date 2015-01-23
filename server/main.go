package main

import (
	"log"
	"net/http"
  "libs/userdb" 
)

const BindAddr string = ":8080"

// check against api keys
// write ip to database (api, ip, old ip)
func main() {


  userdb.Create();

	router := NewRouter()
	log.Printf("Listening on %s", BindAddr)
	log.Fatal(http.ListenAndServe(BindAddr, router))
}

// root iterate apis
// set up iptables, remove old ip, set nil in db
