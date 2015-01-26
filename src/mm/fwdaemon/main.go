package main

import (
	"libs/userdb"
	"log"
	"os/exec"
	"strings"
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

	removeDockRules()
	applyDatabaseRules()

	// every 2s check if database
	doEvery(duration, setupUsersRules)
}

func doEvery(d time.Duration, f func()) {
	for {
		time.Sleep(duration)
		f()
	}
}

// TODO
func removeDockRules() {
	//iptables -D FORWARD -d 172.17.0.3/32 ! -i docker0 -o docker0 -p udp -m udp --dport 53 -j ACCEPT
	//iptables -I INPUT -p udp --dport 53  -j REJECT
}

// TODO clear rules and apply existing from database
func applyDatabaseRules() {
}

func setupUsersRules() {
	users, err := userdb.ChangedUsers()

	if err != nil {
		log.Fatal("Failed to query database", err)
	}

	if len(users) == 0 {
		return
	}

	// TODO parse iptables -nL INPUT
	// TODO should skip existing rules
	for _, user := range users {
		log.Printf("%s changing %s to %s", user.ApiKey, user.OldIp, user.Ip)
		userdb.UpdateUserLastChecked(user)
		cmdIptables("-D", "INPUT", "-s", user.OldIp, "-p", "udp", "-m", "udp", "--dport", "53", "-j", "ACCEPT")
		cmdIptables("-I", "INPUT", "-s", user.Ip, "-p", "udp", "-m", "udp", "--dport", "53", "-j", "ACCEPT")
	}

}

func cmdIptables(args ...string) {
	cmd := exec.Command("iptables", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%s", strings.Join(cmd.Args, " "))
		log.Printf("%s", string(out))
		log.Print(err)
	}
}
