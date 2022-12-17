package app

import (
	"database/sql"
	"embed"
	"fmt"
	"net/http"

	migrate "github.com/rubenv/sql-migrate"
)

//go:embed migrations
var content embed.FS

func Migration() {
	migrations := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(content),
	}

	db, err := sql.Open("sqlite3", "~/.link/link.db")
	if err != nil {
		println(err.Error())
	}

	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		// Handle errors!
		println(err.Error())
	}
	fmt.Printf("Applied %d migrations!\n", n)
}
