# SQL  Communication

Application working with the PostgreSQL database.

## Installing

```
go get github.com/radmirid/sql-communication
```

## Running

```
go run main.go
```

## Usage Example

```go
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
}
```

## LICENSE

[MIT License](LICENSE)
