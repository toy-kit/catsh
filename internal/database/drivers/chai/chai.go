// Package chai defines and registers usql's ChaiSQL driver.
//
// See: https://github.com/chaisql/chai
package chai

import (
	"catsh/internal/database/drivers"

	_ "github.com/chaisql/chai/driver" // DRIVER
)

func init() {
	drivers.Register("chai", drivers.Driver{})
}
