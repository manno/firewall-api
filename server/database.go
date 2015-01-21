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
