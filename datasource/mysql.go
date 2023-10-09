package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tjmtmmnk/mybig/schema"
)

type MySQL struct {
	db *sqlx.DB
}

type Column struct {
	Name     string `db:"COLUMN_NAME"`
	DataType string `db:"DATA_TYPE"`
}

func connect(config ConnectConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
	return sqlx.Open("mysql", dsn)
}

func NewMySQL(config ConnectConfig) MySQL {
	db, err := connect(config)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return MySQL{
		db,
	}
}

func (m MySQL) Columns(tableName string) []schema.Column {
	var mysqlColumns []Column
	sql := fmt.Sprintf("select COLUMN_NAME, DATA_TYPE from INFORMATION_SCHEMA.COLUMNS where table_name = '%s'", tableName)
	err := m.db.Select(&mysqlColumns, sql)
	if err != nil {
		panic(err)
	}
	columns := make([]schema.Column, len(mysqlColumns))
	for i, v := range mysqlColumns {
		columns[i] = schema.Column{
			Name:     v.Name,
			DataType: string(v.DataType),
			Signed:   false,
		}
	}

	return columns
}
