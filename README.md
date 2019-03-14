# 技術書典６用リポジトリ
[技術書典６](https://techbookfest.org/event/tbf06)（2019/4/14）で頒布した書籍で使用するソースが入ったリポジトリです。
書籍の詳細はこちら。
GoでGoogle Cloud Functionsを使う場合の例としてもご使用いただけます。

# セットアップ
- Google Cloud Shell使用
- Google Cloud FunctionsのAPIを有効にする
- Google Cloud Shellで以下のコマンドを実行する

```shell
# gcloud components update && gcloud components install beta
```


# デプロイ方法
## trigger_http
HTTPリクエストで起動するFunctions
```shell
$ gcloud functions deploy TriggerHTTP --runtime go111 --trigger-http
```

## trigger_bucket
Google Cloud Storageのバケットの中にオブジェクトが作成された場合に起動するFunctions

```shell
$ gsutil mb -c nearline gs://YOUR_BUCKET_NAME
$ gcloud functions deploy TriggerStorage --runtime go111 \
  --trigger-resource YOUR_BUCKET_NAME \ 
  --trigger-event google.storage.object.finalize
```

## trigger_pubsub
Google Cloud Pub/Subにメッセージが配信された場合に起動するFunctions

```shell
$ gcloud pubsub topics create YOUR_TOPIC
$ gcloud functions deploy TriggerPubSub --runtime go111 --trigger-topic YOUR_TOPIC
```


## trigger_http_to_bucket
HTTPリクエストで起動して、パラメータをバケットの内のオブジェクトに出力するFunctions

```shell
$ echo "BUCKET_NAME: YOUR_BUCKET_NAME" > .env.yaml
$ gcloud functions deploy TriggerHTTPToBucket --runtime go111 \
     --trigger-http --env-vars-file .env.yaml
```

## trigger_pubsub_to_bigquery  
Google Cloud Pub/Subにメッセージが配信された場合に起動してBigQueryへ出力するFunctions

```shell
$ bq mk GREETINGS
$ bq mk --table GREETINGS.NAMES names.json
$ gcloud functions deploy TriggerPubSubToBigQuery --runtime go111 \
       --trigger-topic YOUR_TOPIC
```

## trigger_pubsub_to_firestore
Google Cloud Pub/Subにメッセージが配信された場合に起動してGoogle Cloud Firestoreへ出力するFunctions

- WebコンソールでFirestoreの管理画面へ行き、ネイティブモードを選択する
- 以下のコマンドでデプロイする
```shell
$ gcloud functions deploy TriggerPubSubToFirestore --runtime go111 \
                   --trigger-topic YOUR_TOPIC
```
