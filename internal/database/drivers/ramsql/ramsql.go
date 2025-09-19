// Package ramsql defines and registers usql's RamSQL driver.
//
// See: https://github.com/proullon/ramsql
package ql

import (
	"catsh/internal/database/drivers"

	_ "github.com/proullon/ramsql/driver" // DRIVER
)

func init() {
	drivers.Register("ramsql", drivers.Driver{})
}
