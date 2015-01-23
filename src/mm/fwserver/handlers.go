package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"libs/models"
	"libs/userdb"
	"net/http"
	"strings"
)

type UserParams struct {
	ApiKey string `json:"api_key"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, models.OkStatus)
}

func Update(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	var userParams UserParams
	if err := json.Unmarshal(body, &userParams); err != nil {
		jsonResponse(w, 422, models.JsonErrorStatus)
		return
	}

	user, err := userdb.FindUser(userParams.ApiKey)
	if err != nil {
		jsonResponse(w, 422, models.InvalidKeyStatus)
		return
	}

	// FIXME parse ip
	user.UpdateIp(extractIp(r.RemoteAddr))
	userdb.UpdateUser(user)

	writeJsonHeader(w, http.StatusOK)
	jsonEncode(w, user)
}

func extractIp(remoteAddr string) string {
	s := strings.Split(remoteAddr, ":")
	if len(s) > 2 {
		// probably ipv6 addr - not supported
		panic(remoteAddr)
	}
	return s[0]
}

func jsonEncode(w http.ResponseWriter, obj interface{}) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		panic(err)
	}
}

func jsonResponse(w http.ResponseWriter, code int, status string) {
	writeJsonHeader(w, code)
	msg := models.Status{State: status}
	jsonEncode(w, msg)
}

func writeJsonHeader(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}
