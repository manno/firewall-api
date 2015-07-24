package main

import (
	"github.com/manno/firewall-api/models"
	"github.com/manno/firewall-api/userdb"
	"log"
	"os/exec"
	"strings"
)

func removeApiChain() {
	cmdIptables("-D", "INPUT", "-p", "udp", "-m", "udp", "--dport", "53", "-j", "FwApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "53", "-j", "FwApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "80", "-j", "FwApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "8123", "-j", "FwApi")
	// move to yaml?
	cmdIptables("-F", "FwApi")
	cmdIptables("-X", "FwApi")
}

func createApiChain() {
	cmdIptables("-N", "FwApi")
	cmdIptables("-I", "INPUT", "-p", "udp", "-m", "udp", "--dport", "53", "-j", "FwApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "53", "-j", "FwApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "80", "-j", "FwApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "8123", "-j", "FwApi")
	// move to yaml (p,port)
	cmdIptables("-A", "FwApi", "-j", "DROP")
}

func initialApplyRules(users models.Users) {
	if len(users) == 0 {
		return
	}

	for _, user := range users {
		cmdIptables("-I", "FwApi", "-s", user.Ip, "-j", "ACCEPT")
	}
}

func applyRules(users models.Users) {
	if len(users) == 0 {
		return
	}

	// TODO parse iptables -nL INPUT
	// TODO should skip existing rules
	for _, user := range users {
		log.Printf("%s changing %s to %s", user.ApiKey, user.OldIp, user.Ip)
		userdb.UpdateUserLastChecked(user)
		cmdIptables("-D", "FwApi", "-s", user.OldIp, "-j", "ACCEPT")
		cmdIptables("-I", "FwApi", "-s", user.Ip, "-j", "ACCEPT")
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
