// Package ots defines and registers usql's Alibaba Tablestore driver.
//
// See: https://github.com/aliyun/aliyun-tablestore-go-sql-driver
package ots

import (
	"catsh/internal/database/drivers"

	_ "github.com/aliyun/aliyun-tablestore-go-sql-driver" // DRIVER
)

func init() {
	drivers.Register("ots", drivers.Driver{})
}
