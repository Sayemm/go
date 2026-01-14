package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
)

func MigrateDB(db *sqlx.DB, dir string) error {
	migrations := &migrate.FileMigrationSource{
		Dir: dir,
	}
	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)
	if err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if n == 0 {
		fmt.Println("Database schema is up to date (no new migrations)")
	} else {
		fmt.Printf("Applied %d migration(s) successfully\n", n)
	}

	return nil
}
