package bigquery

type DataType string

const (
	Bool      DataType = "bool"
	String    DataType = "string"
	Int64     DataType = "int64"
	Numeric   DataType = "numeric"
	DateTime  DataType = "datetime"
	Date      DataType = "Date"
	Timestamp DataType = "timestamp"
	Float64   DataType = "float64"
	Json      DataType = "json"
	Time      DataType = "time"
)

type Column struct {
	Name     string
	DataType DataType
}
