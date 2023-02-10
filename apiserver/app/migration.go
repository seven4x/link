package app

import (
	"database/sql"
	"embed"
	"fmt"
	"os"

	"github.com/seven4x/link/app/log"

	migrate "github.com/rubenv/sql-migrate"
	flag "github.com/spf13/pflag"
)

//go:embed migrations/*
var content embed.FS
var DbPath string

func Migration() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	flag.StringVar(&DbPath, "dbpath", home+"/.link/link.db", "the file path of sqlite db")
	log.Debugf("dbPath is %s", DbPath)

	if _, err := os.Stat(DbPath); os.IsNotExist(err) {
		f, err := os.Create(DbPath)
		if err != nil {
			return err
		}
		f.Close()
	}
	db, err := sql.Open("sqlite3", DbPath)
	if err != nil {
		return err
	}
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: content,
		Root:       "migrations",
	}
	n, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}
	fmt.Printf("Applied %d migrations!\n", n)
	return nil
}
