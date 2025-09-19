// Package bigquery defines and registers usql's Google BigQuery driver.
//
// See: https://github.com/go-gorm/bigquery
package bigquery

import (
	"catsh/internal/database/drivers"

	_ "gorm.io/driver/bigquery/driver" // DRIVER
)

func init() {
	drivers.Register("bigquery", drivers.Driver{})
}
