// Package spanner defines and registers usql's Google Spanner driver.
//
// See: https://github.com/googleapis/go-sql-spanner
package spanner

import (
	"catsh/internal/database/drivers"

	_ "github.com/googleapis/go-sql-spanner" // DRIVER
)

func init() {
	drivers.Register("spanner", drivers.Driver{})
}
