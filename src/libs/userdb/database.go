package userdb

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"libs/models"
	"log"
	"os"
	"time"
)

type UserRow struct {
	Id        int
	ApiKey    string
	Ip        sql.NullString
	OldIp     sql.NullString
	UpdatedAt time.Time
}

const createStmt string = `
  create table users (
    id serial primary key,
    api_key text not null unique,
    ip text,
    old_ip text,
    updated_at timestamp not null,
    last_checked_at timestamp not null
  );
`
const seedStmt = "INSERT INTO users (api_key,updated_at,last_checked_at) VALUES ('123',NOW(),NOW());"
const findByApiKeyStmt = "SELECT api_key, ip, old_ip, updated_at FROM users WHERE api_key = $1"
const updateUserStmt = "UPDATE users SET ip=$1, old_ip=$2, updated_at=$3 WHERE api_key = $4"
const changedUsers = "SELECT api_key, ip, old_ip, updated_at FROM users WHERE updated_at > last_checked_at"
const updateUserLastCheckedStmt = "UPDATE users SET last_checked_at=NOW() WHERE api_key = $1"
const allUsers = "SELECT api_key, ip, old_ip, updated_at FROM users"

var db *sql.DB

func Open() {
	db = connectDatabase()
}

func Close() {
	db.Close()
}

func FindUser(api_key string) (user models.User, err error) {
	rows, err := db.Query(findByApiKeyStmt, api_key)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if !rows.Next() {
		log.Printf("not found, params: (api_key: %s)", api_key)
		return user, errors.New("record not found")
	}
	user, err = scanUserRow(rows)
	if err != nil {
		log.Fatal(err)
	}

	return user, nil
}

func AllUsers() (users models.Users, err error) {
	rows, err := db.Query(allUsers)
	if err != nil {
		log.Printf("%s)", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := scanUserRow(rows)
		if err != nil {
			log.Fatal("Failed to scan")
		}
		users = append(users, user)
	}

	return users, err
}

func ChangedUsers() (users models.Users, err error) {
	rows, err := db.Query(changedUsers)
	if err != nil {
		log.Printf("%s)", err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user, err := scanUserRow(rows)
		if err != nil {
			log.Fatal("Failed to scan")
		}
		users = append(users, user)
	}

	return users, err
}

func UpdateUser(user models.User) {
	stmt, err := db.Prepare(updateUserStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.Ip, user.OldIp, user.UpdatedAt, user.ApiKey); err != nil {
		log.Fatal(err)
	}
}

func UpdateUserLastChecked(user models.User) {
	stmt, err := db.Prepare(updateUserLastCheckedStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(user.ApiKey); err != nil {
		log.Fatal(err)
	}
}

func Create() {
	if _, err := db.Exec(createStmt); err != nil {
		log.Printf("%q: %s\n", err, createStmt)
	}
	if _, err := db.Exec(seedStmt); err != nil {
		log.Printf("%q: %s\n", err, createStmt)
	}
}

func Exists() bool {
	existsStmt := ` 
	SELECT EXISTS (
		SELECT 1 
		FROM   pg_catalog.pg_class c
		JOIN   pg_catalog.pg_namespace n ON n.oid = c.relnamespace
		WHERE  n.nspname = 'public'
		AND    c.relname = 'users'
		AND    c.relkind = 'r'
	);
	`
	var found bool
	err := db.QueryRow(existsStmt).Scan(&found)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return found
}

func convertSqlString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

func connectDatabase() *sql.DB {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("FWDB_USER"),
		os.Getenv("FWDB_PASSWORD"),
		os.Getenv("FWDB_DB"))
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func scanUserRow(rows *sql.Rows) (user models.User, err error) {
	var u UserRow
	if err := rows.Scan(&u.ApiKey, &u.Ip, &u.OldIp, &u.UpdatedAt); err != nil {
		return user, err
	}

	user.ApiKey = u.ApiKey
	user.Ip = convertSqlString(u.Ip)
	user.OldIp = convertSqlString(u.OldIp)
	user.UpdatedAt = u.UpdatedAt
	return user, err
}
