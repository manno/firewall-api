package main

import (
	"libs/userdb"
	"log"
	"time"
)

const duration time.Duration = 200
var lastCheck time.Time = time.Now()

// root iterate apis
// set up iptables, remove old ip, set nil in db
func main() {
	if !userdb.Exists() {
		log.Fatal("Database not found")
	}
	// every 2s check if database 
	doEvery(duration, setupUsers)
}

func doEvery(d time.Duration, f func()) {
	for {
		time.Sleep(duration)
		f()
	}
}

func setupUsers() {
	users,err := userdb.ChangedUsers(lastCheck)
	if err != nil {
		log.Fatal("Failed to query database", err)
	}

	if (len(users) > 0) {
		log.Print(users)
	}

	// select user.* from users where updated_at > last_max
	// set last_max to max(updated_at)
	// iterate users, 
		// if ip!=old_ip
		//   iptables -D old_ip
		//   iptables -A ip
}
