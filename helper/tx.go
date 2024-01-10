package helper

import (
	"database/sql"
)

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback() // rollback if cause error
		IfErrorPanic(errRollback)
		panic(err) // out from function
	} else {
		errorCommit := tx.Commit()
		IfErrorPanic(errorCommit)
	}
}
