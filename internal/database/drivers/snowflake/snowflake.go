// Package snowflake defines and registers usql's Snowflake driver.
//
// See: https://github.com/snowflakedb/gosnowflake
package snowflake

import (
	"io"
	"strconv"

	"catsh/internal/database/drivers"
	"catsh/internal/database/drivers/metadata"
	infos "catsh/internal/database/drivers/metadata/informationschema"

	"catsh/internal/database/env"

	"github.com/snowflakedb/gosnowflake" // DRIVER
	"github.com/xo/tblfmt"
)

func init() {
	gosnowflake.GetLogger().SetOutput(io.Discard)
	newReader := infos.New(
		infos.WithPlaceholder(func(int) string { return "?" }),
		infos.WithCustomClauses(map[infos.ClauseName]string{
			infos.SequenceColumnsIncrement: "''",
		}),
		infos.WithFunctions(false),
		infos.WithIndexes(false),
		infos.WithConstraints(false),
		infos.WithColumnPrivileges(false),
	)
	drivers.Register("snowflake", drivers.Driver{
		AllowMultilineComments: true,
		Err: func(err error) (string, string) {
			if e, ok := err.(*gosnowflake.SnowflakeError); ok {
				return strconv.Itoa(e.Number), e.Message
			}
			return "", err.Error()
		},
		NewMetadataReader: newReader,
		NewMetadataWriter: func(db drivers.DB, w io.Writer, opts ...metadata.ReaderOption) metadata.Writer {
			writerOpts := []metadata.WriterOption{
				metadata.WithListAllDbs(func(pattern string, verbose bool) error {
					return listAllDbs(db, w, pattern, verbose)
				}),
			}
			return metadata.NewDefaultWriter(newReader(db, opts...), writerOpts...)(db, w)
		},
	})
}

func listAllDbs(db drivers.DB, w io.Writer, _ string, _ bool) error {
	rows, err := db.Query("SHOW databases")
	if err != nil {
		return err
	}
	defer rows.Close()
	params := env.Vars().Print()
	params["title"] = "List of databases"
	return tblfmt.EncodeAll(w, rows, params)
}
