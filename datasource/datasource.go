package datasource

import "github.com/tjmtmmnk/mybig/schema"

type SourceType string

const (
	SourceMySQL SourceType = "mysql"
)

type ConnectConfig struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

type DataSource interface {
	Columns(tableName string) []schema.Column
}

func New(sourceType SourceType, config ConnectConfig) DataSource {
	if sourceType == SourceMySQL {
		return NewMySQL(config)
	}
	return NewMySQL(config)
}
