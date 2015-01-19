package main

import "time"

type User struct {
	Id        int       `json:"id"`
	ApiKey    string    `json:"api_key"`
	Ip        string    `json:"ip"`
	OldIp     string    `json:"old_ip"`
	UpdatedAt time.Time `json:"update_at"`
}

type Users []User

const OkStatus string = "ok"

type Status struct {
	State string `json:"status"`
}
