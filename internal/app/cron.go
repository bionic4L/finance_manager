package app

import (
	"context"

	"github.com/jmoiron/sqlx"
)

const queryFillReport = `
INSERT INTO report (from_id, to_id, sum)
SELECT from_id, to_id, SUM(amount) AS sum
FROM transactions
WHERE DATE_TRUNC('month', transactions_date) = DATE_TRUNC('month', CURRENT_DATE)
GROUP BY from_id, to_id;`

func Report(ctx context.Context, db *sqlx.DB) error {
	_, err := db.ExecContext(ctx, queryFillReport)
	if err != nil {
		return err
	}
	return nil
}
