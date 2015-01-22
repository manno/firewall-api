package main

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "os"
  "log"
)

const Repo string = "repo.db"

type Row struct {
	Id        int       `json:"id"`
}

const createStmt string = `
  create table users (
    id integer not null primary key,
    api_key text,
    ip text,
    old_ip text,
    updated_at datetime
  );
`

func FindUser(api_key string) User {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := db.Prepare("SELECT id, api_key, ip, old_ip, updated_at FROM users WHERE api_key = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()
  var user User
  err = stmt.QueryRow(api_key).Scan(&user.Id, &user.ApiKey, &user.Ip, &user.OldIp, &user.UpdatedAt)
  if err != nil {
    log.Fatal(err)
  }
  return user
}

func UpdateUser(user User) {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := db.Prepare("UPDATE users SET (ip, old_ip, update_at) VALUES (?,?,?) WHERE api_key = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()
  _, err = stmt.Exec(user.Ip, user.OldIp, user.UpdatedAt, user.ApiKey)
  if err != nil {
    log.Fatal(err)
  }
}

func RepoCreate() {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  _, err = db.Exec(createStmt)
  if err !=nil {
    log.Printf("%q: %s\n", err, createStmt)
  }
}

func RepoDrop() {
  os.Remove(Repo)
}

func connectDatabase() (*sql.DB, error) {
  return sql.Open("sqlite3", Repo)
}

//func Repo
