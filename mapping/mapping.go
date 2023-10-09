package mapping

import (
	"github.com/tjmtmmnk/mybig/datasource"
	bqType "github.com/tjmtmmnk/mybig/mapping/bigquery"
	myType "github.com/tjmtmmnk/mybig/mapping/mysql"
	"github.com/tjmtmmnk/mybig/schema"
)

func ToBigQueryColumn(sourceType datasource.SourceType, c schema.Column) bqType.Column {
	if sourceType == datasource.SourceMySQL {
		return myType.ToBigQueryColumn(c)
	}
	return bqType.Column{}
}
