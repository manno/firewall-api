package userdb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"libs/models"
	"os"
	"time"
	"log"
)

const RepoFilename string = "repo.db"

type UserRow struct {
	Id        int
	ApiKey    string
	Ip        sql.NullString
	OldIp     sql.NullString
	UpdatedAt time.Time
}

const createStmt string = `
  create table users (
    id integer not null primary key,
    api_key text not null unique,
    ip text,
    old_ip text,
    updated_at datetime not null
  );
`
const seedStmt = "INSERT INTO users (api_key, updated_at) VALUES ('123', date('now'));"
const findByApiKeyStmt = "SELECT id, api_key, ip, old_ip, updated_at FROM users WHERE api_key = ?"
const updateUserStmt = "UPDATE users SET ip=?, old_ip=?, updated_at=? WHERE api_key = ?"

func FindUser(api_key string) (user models.User, err error) {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	stmt, err := db.Prepare(findByApiKeyStmt)
	if err != nil {
		log.Fatal(err)
		return user, err
	}
	defer stmt.Close()

	var u UserRow
	if err = stmt.QueryRow(api_key).Scan(&u.Id, &u.ApiKey, &u.Ip, &u.OldIp, &u.UpdatedAt); err != nil {
		log.Printf("%s params: (api_key: %s)", err, api_key)
		return user, err
	}

	user.ApiKey = u.ApiKey
	user.Ip = convertSqlString(u.Ip)
	user.OldIp = convertSqlString(u.OldIp)
	user.UpdatedAt = u.UpdatedAt
	return user, nil
}

func UpdateUser(user models.User) {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(updateUserStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Ip, user.OldIp, user.UpdatedAt, user.ApiKey); err != nil {
		log.Fatal(err)
	}
}

func Create() {
	db, err := connectDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if _, err = db.Exec(createStmt); err != nil {
		log.Printf("%q: %s\n", err, createStmt)
	}
	if _, err = db.Exec(seedStmt); err != nil {
		log.Printf("%q: %s\n", err, createStmt)
	}
}

func Drop() {
	os.Remove(RepoFilename)
}

func Exists() bool {
	if _, err := os.Stat(RepoFilename); os.IsNotExist(err) {
		return false
	}
	return true
}

func convertSqlString(nullStr sql.NullString) string {
	value, err := nullStr.Value()
	if err == nil {
		return value.(string)
	}
	return ""
}

func connectDatabase() (*sql.DB, error) {
	return sql.Open("sqlite3", RepoFilename)
}
