package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upTransactions, downTransactions)
}

func upTransactions(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS transactions (
				transaction_id SERIAL PRIMARY KEY,
				from_id INT,
				to_id INT,
				FOREIGN KEY (from_id) REFERENCES users(id),
				FOREIGN KEY (to_id) REFERENCES users(id)
);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downTransactions(tx *sql.Tx) error {
	query := `DROP TABLE transactions;`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
