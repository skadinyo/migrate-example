package main

// import (
// 	"log"

// 	"github.com/mattes/migrate"
// 	_ "github.com/mattes/migrate/database/postgres"
// 	_ "github.com/mattes/migrate/source/github"
// )

// func main() {
// 	m, err := migrate.New(
// 		"file://Users/skadinyo/Projects/go/src/github.com/skadinyo/migrate-example/files/database/migrates",
// 		"postgres://skadinyo:kentut@localhost:5432/dsp")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	m.Steps(2)

// }

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func main() {
	db, err := sql.Open("postgres", "postgres://localhost:26257/dsp?sslmode=disable")
	if err != nil {
		log.Println("log 1,", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	if err != nil {
		log.Println("log 2,", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file:////Users/skadinyo/Projects/go/src/github.com/skadinyo/migrate-example/files/database/migrates/",
		"postgres", driver)

	if err != nil {
		log.Println("log 3,", err)
	}

	m.Up()

}
