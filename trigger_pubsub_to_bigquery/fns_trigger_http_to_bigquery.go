package functions

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	//1.BiqQueryを操作するのに必要なライブラリ
	"cloud.google.com/go/bigquery"
)

//2.Pub/Subから受け取るメッセージを格納する構造体
type PubSubMessage struct {
	Data []byte `json:"data"`
}

//3.メッセージの中身を格納し、BigQueryにデータを追加するための構造体、タグで変数とキーを紐づける
type Info struct {
	Name     string    `json:"name" bigquery:"NAME"`
	Place    string    `json:"place" bigquery:"PLACE"`
	Datetime time.Time `bigquery:"DATETIME"`
}

//4.Pub/Subからメッセージを受信した時に実行される
func TriggerPubSubToBigQuery(ctx context.Context, m PubSubMessage) error {
	var i Info

	//6.json形式のメッセージを構造体へ格納する
	err := json.Unmarshal(m.Data, &i)

	//7.エラー時はエラーの型とエラー内容をLoggingへ出力する
	if err != nil {
		log.Printf("メッセージ変換エラー　Error:%T message: %v", err, err)
		return nil
	}

	InsertBigQuery(ctx, i)

	return nil
}

func InsertBigQuery(ctx context.Context, i Info) {

	//5.プロジェクトIDを取得する
	projectID := os.Getenv("GCP_PROJECT")

	//8.BigQueryを操作するクライアントを作成する、エラーの場合はLoggingへ出力する
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("BigQuery接続エラー　Error:%T message: %v", err, err)
		return
	}

	//9.確実にクライアントを閉じるようにする
	defer client.Close()

	//10.クライアントからテーブルを操作するためのアップローダーを取得する
	u := client.Dataset("GREETINGS").Table("NAMES").Uploader()

	//11.現在時刻を構造体へ格納する
	i.Datetime = time.Now()

	items := []Info{i}

	//12.テーブルへデータの追加を行う、エラーの場合はLoggingへ出力する
	err = u.Put(ctx, items)
	if err != nil {
		log.Printf("データ書き込みエラー　Error:%T message: %v", err, err)
		return
	}
}
