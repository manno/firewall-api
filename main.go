package main

//import "encoding/json"
import (
  "fmt"
  "html"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)
//import "os"

// check against api keys
// write ip to database (api, ip, old ip)
func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/ping/{apiKey}", Ping)
  log.Fatal(http.ListenAndServe(":8080", router))
}


func Index(w http.ResponseWriter, r *http.Request) {
  users := Users{
    User{ApiKey: "123", Ip: r.RemoteAddr},
  }
  json.NewEncoder(w).Encode(users)
}

func Ping(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  apiKey := vars["apiKey"]
  fmt.Fprintf(w, "key:", apiKey)
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

// root iterate apis
// set up iptables, remove old ip, set nil in db
