package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upReservations, downReservations)
}

func upReservations(tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS reservations (
				reservation_id SERIAL PRIMARY KEY,
				from_id INT,
				to_id INT,
				amount INT,
				reservation_date TIMESTAMP DEFAULT now(),
				is_pending BOOL DEFAULT TRUE,
				is_done BOOL DEFAULT FALSE,
				is_canceled BOOL DEFAULT FALSE,
				FOREIGN KEY (from_id) REFERENCES users(id),
				FOREIGN KEY (to_id) REFERENCES users(id)
);`
	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func downReservations(tx *sql.Tx) error {
	query := `DROP TABLE reservations;`

	_, err := tx.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
