package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	RegAt    time.Time
}

func main() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres sslmode=disable password=qwerty123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	err = insertUser(db, User{
		Name:     "Bob",
		Email:    "bob@mail.ru",
		Password: "qwerty123",
	})

	users, err := getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	updateUser(db, 2, User{
		Name:  "Sam",
		Email: "sam@mail.ru",
	})

	users, err = getAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, err
}

func getOneUser(db *sql.DB, id int) (User, error) {
	var u User
	err := db.QueryRow("select * from users where id = $1", 2).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegAt)

	return u, err
}

func insertUser(db *sql.DB, u User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3)",
		u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into logs (entity, action) values ($1, $2)",
		"user", "created")
	if err != nil {
		return err
	}

	return tx.Commit()
}

func deleteUser(db *sql.DB, id int) {
	db.Exec("delete from users where id = $1", id)
}

func updateUser(db *sql.DB, id int, updatedUser User) {
	db.Exec("update users set name=$1, email=$2, where id=$3",
		updatedUser.Name, updatedUser.Email, id)
}
