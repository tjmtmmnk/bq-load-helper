package mapping

import (
	bq "github.com/tjmtmmnk/mybig/mapping/bigquery"
	"github.com/tjmtmmnk/mybig/schema"
)

type DataType string

const (
	Bit            DataType = "bit"
	TinyInt        DataType = "tinyint"
	Bool           DataType = "bool"
	Varchar        DataType = "varchar"
	BigInt         DataType = "bigint"
	BigIntUnsigned DataType = "bigint unsigned"
)

func ToBigQueryColumn(c schema.Column) bq.Column {
	var dataType bq.DataType

	switch DataType(c.DataType) {
	case TinyInt:
		dataType = bq.Int64
	case Varchar:
		dataType = bq.String
	case BigInt:
		if c.Signed {
			dataType = bq.Int64
		} else {
			dataType = bq.Numeric
		}
	}

	return bq.Column{
		Name:       c.Name,
		DataType: dataType,
	}
}
