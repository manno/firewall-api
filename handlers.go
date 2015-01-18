package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "html"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
  ok := Status{State: "ok"}
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err:= json.NewEncoder(w).Encode(ok); err != nil {
    panic(err)
  }
}

func Ping(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  apiKey := vars["apiKey"]
  fmt.Fprintf(w, "key:", apiKey)
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
