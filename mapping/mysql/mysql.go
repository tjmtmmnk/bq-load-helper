package mapping

import (
	bqType "github.com/tjmtmmnk/mybig/mapping/bigquery"
	"github.com/tjmtmmnk/mybig/schema"
)

type DataType string

const (
	Bit     DataType = "bit"
	TinyInt DataType = "tinyint"
	Bool    DataType = "bool"
	Varchar DataType = "varchar"
	BigInt  DataType = "bigint"
)

func ToBigQueryColumn(c schema.Column) bqType.Column {
	var dataType bqType.DataType

	columnDataType := DataType(c.DataType)
	switch columnDataType {
	case TinyInt:
		dataType = bqType.Int64
	case Varchar:
		dataType = bqType.String
	case BigInt:
		dataType = bqType.Int64
	}

	return bqType.Column{
		Name:     c.Name,
		DataType: dataType,
	}
}
