package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upDeposits, downDeposits)
}

func upDeposits(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS deposits (
				deposit_id SERIAL PRIMARY KEY NOT NULL,
				user_id INT NOT NULL,
				amount INT NOT NULL,
				transaction_date TIMESTAMP DEFAULT now(),
				FOREIGN KEY (user_id) REFERENCES users(id));`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downDeposits(tx *sql.Tx) error {
	query := `DROP TABLE deposits;`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
