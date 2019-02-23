package functions

import (
	"context"
	"log"

	//1.Pubsubから受け取ったjson形式のメッセージを変換するために必要なライブラリ
	"encoding/json"
)

//2.受け取ったメッセージを格納する構造体、メッセージはBase64でエンコードされている
type PubSubMessage struct {
	Data []byte `json:"data"`
}

//3.メッセージのjsonの中身を格納する構造体、メッセージのjsonの中身を格納する構造体、タグで変数とキーを紐づける
type Info struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

//4.Pub/Subからメッセージを受信した時に実行される
func TriggerPubSub(ctx context.Context, m PubSubMessage) error {
	var i Info

	//5.json形式のメッセージを構造体へ格納する
	err := json.Unmarshal(m.Data, &i)

	//6.エラー時はエラーの型とエラー内容をLoggingへ出力する
	if err != nil {
		log.Printf("Error:%T message: %v", err, err)
		return nil
	}

	////7.メッセージの内容をLoggingへ出力する
	log.Printf("こんにちは、%sさん！%sへCloud Pub/SubからFunctions経由で愛をこめて。", i.Name, i.Place)
	return nil
}
