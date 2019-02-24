package functions

import (
	"fmt"
	"net/http"
)

//1.TriggerHTTP関数がHTTPトリガーで実行される
func TriggerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method { //2.HTTPメソッドにより処理を分岐する。
	case http.MethodGet: //3.GETの場合。
		fmt.Fprint(w, "Cloud Functions より Hello World!（GETメソッド）")
	case http.MethodPost: //4.POSTの場合。
		fmt.Fprint(w, "Cloud Functions より Hello World!（POSTメソッド）")
	default: //5.それ以外の場合はエラー
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
