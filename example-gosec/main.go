package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
)

const MySQLUser = "example"
const MySQLPassword = "zVxqYh2v0BZFVMBSUv6Kg9yCswGrYcZT"
const MySQLHost = "127.0.0.1"
const MySQLDatabase = "test"

type User struct {
	Username     string
	PasswordHash string
}

func NewUser(username string) *User {
	pwd := make([]byte, 12)
	rand.Read(pwd)

	u := User{
		Username:     username,
		PasswordHash: fmt.Sprintf("%x", md5.New().Sum(pwd)),
	}

	return &u
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", MySQLUser, MySQLPassword, MySQLHost, MySQLDatabase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	user := NewUser("martin")
	db.Query("INSERT INTO users (username, password_hash), VALUES ('" + user.Username + "', '" + user.PasswordHash + "')")
}
