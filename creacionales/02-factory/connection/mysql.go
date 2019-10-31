package connection

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	db *sql.DB
}

func (m *MySql) Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?tls=false&autocommit=true&allowNativePasswords=true&parseTime=true",
		"root",
		"root",
		"127.0.0.1",
		"3306",
		"mysql",
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	m.db = db
	return nil
}

func (m *MySql) GetNow() (*time.Time, error) {
	t := &time.Time{}
	err := m.db.QueryRow("select CURDATE()").Scan(t)
	if err != nil {
		fmt.Printf("error al leer la fecha del servidor: %v", err)
		return nil, err
	}

	return t, nil
}

func (m *MySql) Close() error {
	return m.db.Close()
}