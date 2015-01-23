package userdb

import (
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
  "os"
  "log"
  "libs/models"
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

func FindUser(api_key string) models.User {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := db.Prepare("SELECT id, api_key, ip, old_ip, updated_at FROM users WHERE api_key = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()
  var user models.User
  if err = stmt.QueryRow(api_key).Scan(&user.Id, &user.ApiKey, &user.Ip, &user.OldIp, &user.UpdatedAt); err != nil {
    log.Fatal(err)
  }
  return user
}

func UpdateUser(user models.User) {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  stmt, err := db.Prepare("UPDATE users SET (ip, old_ip, update_at) VALUES (?,?,?) WHERE api_key = ?")
  if err != nil {
    log.Fatal(err)
  }
  defer stmt.Close()

  if _, err := stmt.Exec(user.Ip, user.OldIp, user.UpdatedAt, user.ApiKey); err !=nil {
    log.Fatal(err)
  }
}

func Create() {
  db, err := connectDatabase()
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  if _, err = db.Exec(createStmt); err !=nil {
    log.Printf("%q: %s\n", err, createStmt)
  }
}

func Drop() {
  os.Remove(Repo)
}

func connectDatabase() (*sql.DB, error) {
  return sql.Open("sqlite3", Repo)
}

//func Repo
