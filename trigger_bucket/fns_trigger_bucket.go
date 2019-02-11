package functions

import (
	"context"
	"fmt"
	"log"
	"time"

	//1.トリガーしたイベントの詳細情報を取得するためのライブラリ
	"cloud.google.com/go/functions/metadata"
)

//2.Google Cloud Storageからトリガーされたイベントを拾うための構造体
type GCSEvent struct {
	Bucket         string    `json:"bucket"` //3.各属性にはjsonタグが必要
	Name           string    `json:"name"`
	Metageneration string    `json:"metageneration"`
	ResourceState  string    `json:"resourceState"`
	TimeCreated    time.Time `json:"timeCreated"`
	Updated        time.Time `json:"updated"`
}

//3.TriggetStorage関数がStorage（バケット）にファイルが作成されたタイミングで実行される
func TriggetStorage(ctx context.Context, e GCSEvent) error {
	//4.トリガーしたイベントの詳細情報を取り出す
	meta, err := metadata.FromContext(ctx)
	if err != nil {
		return fmt.Errorf("metadata.FromContext: %v", err)
	}
	//5.4で取り出した詳細情報と構造体に格納されたイベントの情報をLoggingに出力する
	log.Printf("Event ID: %v\n", meta.EventID)
	log.Printf("Event type: %v\n", meta.EventType)
	log.Printf("バケット名: %v\n", e.Bucket)
	log.Printf("ファイル名: %v\n", e.Name)
	log.Printf("リソース状態: %v\n", e.ResourceState)
	log.Printf("作成日時: %v\n", e.TimeCreated)
	log.Printf("更新日時: %v\n", e.Updated)
	return nil
}
