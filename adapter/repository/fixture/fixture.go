package fixture

import (
	"io/fs"
	"database/sql"
	"log"
	"maragu.dev/migrate"
	_ "github.com/mattn/go-sqlite3"
	"context"
)


func Up(migrationsDir fs.FS) *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	if err := migrate.Up(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	return db
}

func Down(db *sql.DB, migrationsDir fs.FS) {
	if err := migrate.Down(context.Background(), db, migrationsDir); err != nil {
		panic(err)
	}
	db.Close()
}