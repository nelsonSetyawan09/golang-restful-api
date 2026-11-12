package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if recovered := recover(); recovered != nil {
		if err := tx.Rollback(); err != nil {
			panic(err)
		}
		panic(recovered)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
