package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func mapDate(d time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{Time: d, Valid: true}
}
