// Package cosmos defines and registers usql's Azure CosmosDB driver.
//
// See: https://github.com/btnguyen2k/gocosmos
package cosmos

import (
	"catsh/internal/database/drivers"

	_ "github.com/btnguyen2k/gocosmos" // DRIVER
)

func init() {
	drivers.Register("cosmos", drivers.Driver{
		Process: drivers.StripTrailingSemicolon,
	}, "gocosmos")
}
