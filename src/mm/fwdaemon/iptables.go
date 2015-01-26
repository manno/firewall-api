package main

import (
	"libs/models"
	"libs/userdb"
	"log"
	"os/exec"
	"strings"
)

func removeApiChain() {
	cmdIptables("-D", "INPUT", "-p", "udp", "-m", "udp", "--dport", "53", "-j", "DnsApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "53", "-j", "DnsApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "80", "-j", "DnsApi")
	cmdIptables("-D", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "443", "-j", "DnsApi")
	cmdIptables("-F", "DnsApi")
	cmdIptables("-X", "DnsApi")
}

func createApiChain() {
	cmdIptables("-N", "DnsApi")
	cmdIptables("-I", "INPUT", "-p", "udp", "-m", "udp", "--dport", "53", "-j", "DnsApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "53", "-j", "DnsApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "80", "-j", "DnsApi")
	cmdIptables("-I", "INPUT", "-p", "tcp", "-m", "tcp", "--dport", "443", "-j", "DnsApi")
	cmdIptables("-A", "DnsApi", "-j", "DROP")
}

func initialApplyRules(users models.Users) {
	if len(users) == 0 {
		return
	}

	for _, user := range users {
		cmdIptables("-I", "DnsApi", "-s", user.Ip, "-j", "ACCEPT")
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
		cmdIptables("-D", "DnsApi", "-s", user.OldIp, "-j", "ACCEPT")
		cmdIptables("-I", "DnsApi", "-s", user.Ip, "-j", "ACCEPT")
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
