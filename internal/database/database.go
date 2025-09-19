package database

import (
	"catsh/global"
	"catsh/internal/database/drivers"
	"catsh/internal/database/internal"
	"catsh/types"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"os"
	"os/user"
	"path/filepath"

	"github.com/dromara/dongle"
	"github.com/dromara/dongle/crypto/cipher"
	"github.com/google/uuid"
	"github.com/xo/dburl"
)

type Conn struct {
	uri      *dburl.URL
	db       *sql.DB
	database types.Database
}

var (
	AesKey    = "dongle1234567890"
	AesIV     = "1234567890123456"
	databases = make(map[string]types.Database)
	Conns     = make(map[string]Conn)
)

func haveConn(id string) (Conn, error) {
	c, ok := Conns[id]
	if ok {
		return c, nil
	}
	database, ok := databases[id]
	if !ok {
		return Conn{}, errors.New("no such database")
	}
	u, err := dburl.Parse(database.URL())
	if err != nil {
		return Conn{}, err
	}
	conn, err := drivers.Open(context.Background(), u, nil, nil)
	if err != nil {
		return Conn{}, err
	}
	Conns[id] = Conn{database: database, db: conn, uri: u}
	return Conns[id], nil
}

func Close(id string) error {
	c, ok := Conns[id]
	if !ok {
		return nil
	}
	defer delete(Conns, id)
	return c.db.Close()
}

func Query(id string, query string) ([]string, any, error) {
	c, err := haveConn(id)
	if err != nil {
		return nil, nil, err
	}
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()
	columns, err := drivers.Columns(c.uri, rows)
	if err != nil {
		return nil, nil, err
	}
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}
	var result []map[string]any
	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			return nil, nil, err
		}
		rowMap := make(map[string]any)
		for i, col := range columns {
			rowMap[col] = values[i]
		}
		result = append(result, rowMap)
	}
	if err = rows.Err(); err != nil {
		return nil, nil, err
	}
	return columns, result, nil
}

func GetDrivers() map[string]string {
	return internal.KnownBuildTags()
}

func GetDatabase() (map[string]types.Database, error) {
	items := make(map[string]types.Database)
	dbDir := getDbDir()
	files, err := os.ReadDir(dbDir)
	if err != nil {
		return items, err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(dbDir, file.Name()))
		if err != nil {
			continue
		}
		var item types.Database
		err = json.Unmarshal(b, &item)
		if err != nil {
			continue
		}
		items[item.Id] = item
	}
	databases = items
	return items, nil
}
func SaveDatabase(db types.Database) error {
	if db.Id == "" {
		db.Id = uuid.New().String()
	}
	b, err := json.Marshal(db)
	if err != nil {
		return err
	}
	c := cipher.NewAesCipher(cipher.CBC)
	c.SetKey([]byte(AesKey))
	c.SetIV([]byte(AesIV))
	c.SetPadding(cipher.PKCS7)
	encrypter := dongle.Encrypt.FromBytes(b).ByAes(c)
	if encrypter.Error != nil {
		return encrypter.Error
	}
	return os.WriteFile(filepath.Join(getDbDir(), db.Id), encrypter.ToBase64Bytes(), 0644)
}

func getDbDir() string {
	dbDir := filepath.Join("."+global.AppConfig.Info.CompanyName, global.AppConfig.Info.ProductName, "db")
	dir := ""
	if u, err := user.Current(); err == nil {
		dir = u.HomeDir
	}
	dbDir = filepath.Join(dir, dbDir)
	os.MkdirAll(dbDir, os.ModePerm)
	return dbDir
}
