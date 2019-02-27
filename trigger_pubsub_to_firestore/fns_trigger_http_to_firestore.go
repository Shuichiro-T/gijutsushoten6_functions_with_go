package functions

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	//1.Firestoreを操作するのに必要なライブラリ
	"cloud.google.com/go/firestore"
)

//2.Pub/Subから受け取るメッセージを格納する構造体
type PubSubMessage struct {
	Data []byte `json:"data"`
}

//3.メッセージの中身を格納し、Firestoreにデータを追加するための構造体、タグで変数とキーを紐づける
type Info struct {
	Name     string    `json:"name" firestore::"NAME"`
	Place    string    `json:"place" firestore:"PLACE"`
	Datetime time.Time `firestore:"DATETIME"`
}

//4.Pub/Subからメッセージを受信した時に実行される
func TriggerPubSubToFirestore(ctx context.Context, m PubSubMessage) error {
	var i Info

	//5.json形式のメッセージを構造体へ格納する
	err := json.Unmarshal(m.Data, &i)

	//6.エラー時はエラーの型とエラー内容をLoggingへ出力する
	if err != nil {
		log.Printf("メッセージ変換エラー　Error:%T message: %v", err, err)
		return nil
	}

	//7.Firestoreにデータを追加する関数を呼び出す
	PutFirestore(ctx, i)

	return nil
}

//8.Firestoreにデータを追加する関数を定義する
func PutFirestore(ctx context.Context, i Info) {

	//9.プロジェクトIDを取得する
	projectID := os.Getenv("GCP_PROJECT")

	//10.Firestoreを操作するクライアントを作成する、エラーの場合はLoggingへ出力する
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Printf("Firestore接続エラー　Error:%T message: %v", err, err)
		return
	}

	//11.確実にクライアントを閉じるようにする
	defer client.Close()

	//12.現在時刻を構造体へ格納する
	i.Datetime = time.Now()

	//13.Firestoreの追加を行う、エラーの場合はLoggingへ出力する
	_, _, err = client.Collection("NAMES").Add(ctx, i)
	if err != nil {
		log.Printf("データ書き込みエラー　Error:%T message: %v", err, err)
		return
	}
}
