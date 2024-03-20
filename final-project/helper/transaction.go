package helper

import (
	"database/sql"
	"net/http"
)

func RollbackOrCommit(tx *sql.Tx, err *error) {
	if *err != nil {
		rollBackError := tx.Rollback()
		if rollBackError != nil {
			*err = NewResponseError(ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	commitError := tx.Commit()
	if commitError != nil {
		*err = NewResponseError(ErrInternal, http.StatusInternalServerError)
	}

}
