package main

import (
	"fmt"
	"github.com/tjmtmmnk/mybig/datasource"
	"github.com/tjmtmmnk/mybig/output"
)

func main() {
	config := datasource.ConnectConfig{
		Host:     "127.0.0.1",
		Port:     "13306",
		DBName:   "sample",
		User:     "root",
		Password: "",
	}
	source := datasource.New(datasource.SourceMySQL, config)
	f := output.Format(source, "test")
	fmt.Println(f.Schema)
	fmt.Println(f.Data)
}

/*
1. テーブル名を指定してcolumnsを返す。このときMySQLの型定義なのかPostgreSQLの型定義なのかを見分けられるようにする
2. 型定義からDatastreamの型対応表を使ってBigQueryでの型との対応を返す
3.[カラム名]:[型],[カラム名]:[型],... というスキーマファイルを吐き出す
4. テーブルから取得したデータからcsvファイルを作る
5. bq load コマンドでcsvファイルとスキーマファイルからデータを取り込む
*/

/*
責務としては…
- datasource: データを集めてくるだけ。データソースの種類とそのデータをカラムの型付きで返してくれる
- mapping: datasourceで集めたデータの型をBigQueryの型に変換する
- output: BigQueryで取り込めるようにクエリを作る
*/
