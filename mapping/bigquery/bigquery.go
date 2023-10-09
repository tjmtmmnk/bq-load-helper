package bigquery

type ColumnType string

const (
	Bool    ColumnType = "bool"
	String  ColumnType = "string"
	Int64   ColumnType = "int64"
	Numeric ColumnType = "numeric"
)

type Column struct {
	Name       string
	ColumnType ColumnType
}
