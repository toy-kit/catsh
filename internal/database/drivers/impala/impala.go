// Package impala defines and registers usql's Apache Impala driver.
//
// See: https://github.com/sclgo/impala-go
package impala

import (
	"context"
	"database/sql"
	"errors"

	"catsh/internal/database/drivers"
	meta "catsh/internal/database/drivers/metadata/impala"

	"github.com/sclgo/impala-go" // DRIVER
)

func init() {
	drivers.Register("impala", drivers.Driver{
		NewMetadataReader: meta.New,
		Copy: func(ctx context.Context, db *sql.DB, rows *sql.Rows, table string) (int64, error) {
			placeholder := func(int) string {
				return "?"
			}
			return drivers.FlexibleCopyWithInsert(ctx, db, rows, table, placeholder, false)
		},
		IsPasswordErr: func(err error) bool {
			var authError *impala.AuthError
			return errors.As(err, &authError)
		},
	})
}
