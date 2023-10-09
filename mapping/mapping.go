package mapping

import (
	"github.com/tjmtmmnk/mybig/datasource"
	bq "github.com/tjmtmmnk/mybig/mapping/bigquery"
	ms "github.com/tjmtmmnk/mybig/mapping/mysql"
	"github.com/tjmtmmnk/mybig/schema"
)

func ToBigQueryColumn(sourceType datasource.SourceType, c schema.Column) bq.Column {
	if sourceType == datasource.SourceMySQL {
		return ms.ToBigQueryColumn(c)
	}
	return bq.Column{}
}
