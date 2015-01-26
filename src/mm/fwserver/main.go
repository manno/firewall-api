package main

import (
	"libs/userdb"
	"log"
	"net/http"
  "os"
)

// check against api keys
// write ip to database (api, ip, old ip)
func main() {
	userdb.Open()
	defer userdb.Close()

	if !userdb.Exists() {
		userdb.Create()
		log.Printf("Database created")
	}

	router := NewRouter()
  var port = ":"+os.Getenv("FW_PORT")
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
