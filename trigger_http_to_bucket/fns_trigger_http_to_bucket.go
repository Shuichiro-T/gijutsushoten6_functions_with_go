package functions

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

//1.TriggerHTTPToBucket関数がHTTPトリガーで実行される
func TriggerHTTPToBucket(w http.ResponseWriter, r *http.Request) {
	switch r.Method { //2.HTTPメソッドにより処理を分岐する
	case http.MethodGet: //3.GETの場合

		//4.GETパラメータの"name"から値を取り出す
		names, err := r.URL.Query()["name"]

		//5.取り出せない場合はエラーとして処理を終了する
		if !err || len(names[0]) < 1 {
			fmt.Fprint(w, "パラメータに\"name\"がありません。\r\n")
			return
		}

		//Storageへ出力する関数を呼び出す
		MakeFile(w, names[0])

	default: //6.それ以外の場合はエラー
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func MakeFile(rw http.ResponseWriter, name string) {

	//7.環境変数からバケット名を取得する
	bucketName := os.Getenv("BUCKET_NAME")

	//8.Storageへ接続するクライアントを初期化する
	ctx := context.Background()
	client, err := storage.NewClient(ctx)

	//9.エラーの場合は処理を終了する
	if err != nil {
		fmt.Fprint(rw, "Storage接続エラー　エラータイプ：%T、エラーメッセージ：%s", err, err)
		return

	}
	//10.接続したクライアントは確実に切断させる
	defer client.Close()

	//11.オブジェクト名を決める yyyyMMddhhmmss形式
	objectName := time.Now().Format("20060102150405")

	//12.オブジェクトを作成して書き込みストリームを開く
	fw := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)

	//13.ファイルを書き込み　エラーの場合は処理を終了する
	if _, err := fw.Write([]byte(name + "\r\n")); err != nil {
		fmt.Fprint(rw, "オブジェクト書き込みエラー　エラータイプ：%T、エラーメッセージ：%s", err, err)
		return
	}

	//14.書き込みストリームを閉じる、エラーの場合は処理を終了する
	if err := fw.Close(); err != nil {
		fmt.Fprint(rw, "オブジェクト切断エラー　エラータイプ：%T、エラーメッセージ：%s", err, err)
		return
	}
}
