package mysql

import (
	bq "github.com/tjmtmmnk/bq-load-helper/mapping/bigquery"
	"github.com/tjmtmmnk/bq-load-helper/schema"
)

type DataType string

const (
	BigInt     DataType = "bigint"
	Binary     DataType = "binary"
	Bit        DataType = "bit"
	Blob       DataType = "blob"
	Char       DataType = "char"
	Date       DataType = "date"
	DateTime   DataType = "datetime"
	Decimal    DataType = "decimal"
	Double     DataType = "double"
	Enum       DataType = "enum"
	Float      DataType = "float"
	Int        DataType = "int"
	Json       DataType = "json"
	LongBlob   DataType = "longblob"
	LongText   DataType = "longtext"
	MediumBlob DataType = "mediumblob"
	MediumInt  DataType = "mediumint"
	MediumText DataType = "mediumtext"
	Set        DataType = "set"
	SmallInt   DataType = "smallint"
	Text       DataType = "text"
	Time       DataType = "time"
	Timestamp  DataType = "timestamp"
	TinyBlob   DataType = "tinyblob"
	TinyInt    DataType = "tinyint"
	TinyText   DataType = "tinytext"
	VarBinary  DataType = "varbinary"
	VarChar    DataType = "varchar"
	Year       DataType = "year"
)

func ToBigQueryColumn(c schema.Column) bq.Column {
	var dataType bq.DataType

	switch DataType(c.DataType) {
	case Bit, Int, MediumInt, SmallInt, TinyInt, Year:
		dataType = bq.Int64
	case Binary, Blob, Char, Enum, LongBlob, LongText, MediumBlob, MediumText, Set, Text, TinyBlob, TinyText, VarBinary, VarChar:
		dataType = bq.String
	case BigInt:
		if c.Signed {
			dataType = bq.Int64
		} else {
			dataType = bq.Numeric
		}
	case Date:
		dataType = bq.Date
	case DateTime:
		dataType = bq.DateTime
	case Decimal:
		dataType = bq.Numeric
	case Double, Float:
		dataType = bq.Float64
	case Json:
		dataType = bq.Json
	case Time:
		dataType = bq.Time
	case Timestamp:
		dataType = bq.Timestamp
	}

	return bq.Column{
		Name:     c.Name,
		DataType: dataType,
	}
}
