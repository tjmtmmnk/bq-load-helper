package mapping

import (
	"github.com/tjmtmmnk/bq-load-helper/datasource"
	bq "github.com/tjmtmmnk/bq-load-helper/mapping/bigquery"
	ms "github.com/tjmtmmnk/bq-load-helper/mapping/mysql"
	"github.com/tjmtmmnk/bq-load-helper/schema"
)

func ToBigQueryColumn(sourceType datasource.SourceType, c schema.Column) bq.Column {
	if sourceType == datasource.SourceMySQL {
		return ms.ToBigQueryColumn(c)
	}
	return bq.Column{}
}
