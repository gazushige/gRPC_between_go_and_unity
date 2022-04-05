# Go言語サーバーとUnityクライアントの接続テストを手軽に実装

## 使い方

### サーバーサイド

```sh
go run cmd/main.go
```

で実行する。Now listeningと表示されれば起動は成功。

### クライアントサイド

1. Unity新規プロジェクトに[unity_gRPC_package](https://packages.grpc.io/)をインストールする。
2. このリポジトリにあるUnity_gRPC_protobuf_Testをインストールする。
3. controller.csを適当な空オブジェクトにアタッチして実行する。
4. ログが返信されれば成功。
