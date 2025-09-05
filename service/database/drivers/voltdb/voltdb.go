// Package voltdb defines and registers usql's VoltDB driver.
//
// See: https://github.com/VoltDB/voltdb-client-go
package voltdb

import (
	"catsh/service/database/drivers"

	_ "github.com/VoltDB/voltdb-client-go/voltdbclient" // DRIVER
)

func init() {
	drivers.Register("voltdb", drivers.Driver{
		AllowMultilineComments: true,
	})
}
