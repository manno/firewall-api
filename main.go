package main

import (
  "log"
  "net/http"
)

type Config struct {
  BindAddr  string
}

// check against api keys
// write ip to database (api, ip, old ip)
func main() {
  router := NewRouter()
  config := Config{BindAddr: ":8080"}
  log.Printf("Listening on %s", config.BindAddr)
  log.Fatal(http.ListenAndServe(config.BindAddr, router))
}

// root iterate apis
// set up iptables, remove old ip, set nil in db
