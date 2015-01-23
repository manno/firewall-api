package models

import "time"

type User struct {
	ApiKey    string    `json:"api_key"`
	Ip        string    `json:"ip"`
	OldIp     string    `json:"old_ip"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Users []User

func (user *User) UpdateIp(ip string) {
	user.OldIp = user.Ip
	user.Ip = ip
	user.UpdatedAt = time.Now().UTC()
}
