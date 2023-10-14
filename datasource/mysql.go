package datasource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tjmtmmnk/bq-load-helper/schema"
	"regexp"
)

type MySQL struct {
	db *sqlx.DB
}

type Column struct {
	Name       string `db:"COLUMN_NAME"`
	DataType   string `db:"DATA_TYPE"`
	ColumnType string `db:"COLUMN_TYPE"`
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

func (m MySQL) Type() SourceType {
	return SourceMySQL
}

func (m MySQL) Query(tableName string) [][]string {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	var results [][]string

	for rows.Next() {
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))

		for i := range values {
			pointers[i] = &values[i]
		}

		if err := rows.Scan(pointers...); err != nil {
			return nil
		}

		stringValues := make([]string, len(columns))
		for i, v := range values {
			switch data := v.(type) {
			case []byte:
				stringValues[i] = string(data)
			case string:
				stringValues[i] = data
			default:
				if data == nil {
					stringValues[i] = ""
				} else {
					stringValues[i] = fmt.Sprint(data)
				}
			}
		}

		results = append(results, stringValues)
	}

	return results
}

func (m MySQL) Columns(tableName string) []schema.Column {
	var mysqlColumns []Column
	sql := fmt.Sprintf("select COLUMN_NAME, DATA_TYPE, COLUMN_TYPE from INFORMATION_SCHEMA.COLUMNS where table_name = '%s'", tableName)
	err := m.db.Select(&mysqlColumns, sql)
	if err != nil {
		panic(err)
	}
	columns := make([]schema.Column, len(mysqlColumns))
	unsignedRe := regexp.MustCompile(`unsigned`)
	for i, v := range mysqlColumns {
		unsigned := unsignedRe.MatchString(v.ColumnType)
		columns[i] = schema.Column{
			Name:     v.Name,
			DataType: v.DataType,
			Signed:   !unsigned,
		}
	}

	return columns
}
