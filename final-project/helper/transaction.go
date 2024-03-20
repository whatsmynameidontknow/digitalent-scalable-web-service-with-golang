package helper

import (
	"database/sql"
	"log/slog"
	"net/http"
)

func RollbackOrCommit(tx *sql.Tx, err *error, log *slog.Logger) {
	if *err != nil {
		rollBackError := tx.Rollback()
		if rollBackError != nil {
			log.Error(rollBackError.Error(), "cause", "tx.Rollback")
			*err = NewResponseError(ErrInternal, http.StatusInternalServerError)
		}
		return
	}

	commitError := tx.Commit()
	if commitError != nil {
		log.Error(commitError.Error(), "cause", "tx.Commit")
		*err = NewResponseError(ErrInternal, http.StatusInternalServerError)
	}
}
