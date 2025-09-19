// Package maxcompute defines and registers usql's Alibaba MaxCompute driver.
//
// See: https://github.com/sql-machine-learning/gomaxcompute
package maxcompute

import (
	"catsh/internal/database/drivers"

	_ "sqlflow.org/gomaxcompute" // DRIVER
)

func init() {
	drivers.Register("maxcompute", drivers.Driver{})
}
