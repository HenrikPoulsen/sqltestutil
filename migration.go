package sqltestutil

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations reads all of the files matching *.up.sql in migrationDir and
// executes them in lexicographical order against the provided db. A typical
// convention is to use a numeric prefix for each new migration, e.g.:
//
//	001_create_users.up.sql
//	002_create_posts.up.sql
//	003_create_comments.up.sql
//
// Note that this function does not check whether the migration has already been
// run. Its primary purpose is to initialize a test database.
func RunMigrations(db *sql.DB, migrationDir string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		migrationDir,
		"postgres", driver)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	return nil
}
