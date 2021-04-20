package sqlw

import (
	"context"
	"database/sql"
)

// Tx calls function f in a transaction.
//
// If f returns error or panic, Tx will roll back the transaction.
func Tx(ctx context.Context, db *sql.DB, opts *sql.TxOptions, f func(tx *sql.Tx) error) (err error) {
	var tx *sql.Tx
	tx, err = db.BeginTx(ctx, opts)
	if err != nil {
		return
	}

	fPanic := true
	defer func() {
		if err != nil || fPanic {
			e := tx.Rollback()
			if err == nil {
				err = e
			}
		} else {
			err = tx.Commit()
		}
	}()
	err = f(tx)
	fPanic = false
	return
}
