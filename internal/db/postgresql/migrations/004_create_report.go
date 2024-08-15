package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upReports, downReports)
}

func upReports(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS report (
				from_id INT,
				to_id INT,
				sum INT);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downReports(tx *sql.Tx) error {
	query := `DROP TABLE report;`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
