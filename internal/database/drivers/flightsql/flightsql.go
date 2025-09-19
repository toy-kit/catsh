// Package flightsql defines and registers usql's FlightSQL driver.
//
// See: https://github.com/apache/arrow/tree/main/go/arrow/flight/flightsql/driver
package flightsql

import (
	"catsh/internal/database/drivers"

	_ "github.com/apache/arrow/go/v17/arrow/flight/flightsql/driver" // DRIVER
)

func init() {
	drivers.Register("flightsql", drivers.Driver{})
}
