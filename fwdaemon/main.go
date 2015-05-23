package main

import (
	"github.com/manno/firewall-api/userdb"
	"log"
	"time"
)

const duration time.Duration = 2000 * time.Millisecond

// set up iptables, remove old ip, set nil in db
func main() {
	userdb.Open()
	defer userdb.Close()

	if !userdb.Exists() {
		userdb.Create()
		log.Printf("Database created")
	}

	removeApiChain()
	createApiChain()
	setupDatabaseRules()

	// every 2s check if database
	doEvery(duration, setupUsersRules)
}

func doEvery(d time.Duration, f func()) {
	for {
		time.Sleep(duration)
		f()
	}
}

// clear rules and apply existing from database
func setupDatabaseRules() {
	users, err := userdb.AllUsers()

	if err != nil {
		log.Fatal("Failed to query database", err)
	}

	initialApplyRules(users)
}

func setupUsersRules() {
	users, err := userdb.ChangedUsers()

	if err != nil {
		log.Fatal("Failed to query database", err)
	}

	applyRules(users)
}
