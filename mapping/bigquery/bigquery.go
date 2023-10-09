package bigquery

type DataType string

const (
	Bool    DataType = "bool"
	String  DataType = "string"
	Int64   DataType = "int64"
	Numeric DataType = "numeric"
)

type Column struct {
	Name       string
	DataType DataType
}
