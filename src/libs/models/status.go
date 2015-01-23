package models

const OkStatus string = "ok"
const JsonErrorStatus string = "json invalid"
const InvalidKeyStatus string = "api key invalid"

type Status struct {
	State string `json:"status"`
}
