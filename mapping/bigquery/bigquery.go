package bigquery

type DataType string

const (
	Bool   DataType = "bool"
	String DataType = "string"
	Int64  DataType = "int64"
)

type Column struct {
	Name     string
	DataType DataType
}
