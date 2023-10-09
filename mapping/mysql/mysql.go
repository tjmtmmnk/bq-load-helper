package mapping

import (
	bqType "github.com/tjmtmmnk/mybig/mapping/bigquery"
	"github.com/tjmtmmnk/mybig/schema"
)

type ColumnType string

const (
	Bit            ColumnType = "bit"
	TinyInt        ColumnType = "tinyint"
	Bool           ColumnType = "bool"
	Varchar        ColumnType = "varchar"
	BigInt         ColumnType = "bigint"
	BigIntUnsigned ColumnType = "bigint unsigned"
)

func ToBigQueryColumn(c schema.Column) bqType.Column {
	var columnType bqType.ColumnType

	switch ColumnType(c.ColumnType) {
	case TinyInt:
		columnType = bqType.Int64
	case Varchar:
		columnType = bqType.String
	case BigInt:
		columnType = bqType.Int64
	case BigIntUnsigned:
		columnType = bqType.Numeric
	}

	return bqType.Column{
		Name:       c.Name,
		ColumnType: columnType,
	}
}
