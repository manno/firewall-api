package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
  "libs/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ok := models.Status{State: models.OkStatus}
	writeJsonHeader(w, http.StatusOK)
	jsonEncode(w, ok)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048567))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	type UserParams struct {
		ApiKey string `json:"api_key"`
	}
	var userParams UserParams

	if err := json.Unmarshal(body, &userParams); err != nil {
		writeJsonHeader(w, 422)
		jsonEncode(w, err)
		return
	}

	var user models.User
	writeJsonHeader(w, http.StatusOK)
	jsonEncode(w, user)
}

func jsonEncode(w http.ResponseWriter, obj interface{}) {
	if err := json.NewEncoder(w).Encode(obj); err != nil {
		panic(err)
	}
}

func writeJsonHeader(w http.ResponseWriter, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
}
