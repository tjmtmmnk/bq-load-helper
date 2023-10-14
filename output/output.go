package output

import (
	"github.com/tjmtmmnk/bq-load-helper/datasource"
	"github.com/tjmtmmnk/bq-load-helper/mapping"
	"strings"
)

type BqLoadFormat struct {
	Schema string
	Data   string
}

func Format(source datasource.DataSource, tableName string) BqLoadFormat {
	columns := source.Columns(tableName)
	schemaRows := make([]string, len(columns))
	for i, c := range columns {
		bqColumn := mapping.ToBigQueryColumn(source.Type(), c)
		schemaRows[i] = bqColumn.Name + ":" + string(bqColumn.DataType)
	}

	rows := source.Query(tableName)
	dataRows := make([]string, len(rows))
	for i, row := range rows {
		dataRows[i] = strings.Join(row, ",")
	}
	//dataRows := make([]string, len(rows))
	return BqLoadFormat{
		Schema: strings.Join(schemaRows, ","),
		Data:   strings.Join(dataRows, "\n"),
	}
}
