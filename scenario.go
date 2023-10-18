package sqltestutil

import (
	"database/sql"
	"os"
)

// LoadScenario reads a SQL "scenario" file and uses it to populate the given
// db.
func LoadScenario(db *sql.DB, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}

	return nil
}
