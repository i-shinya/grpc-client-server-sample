# grpc-client-server-sample

golang で grpc サンプルを実装するリポジトリ
クライアントとサーバの実装サンプルを実装していく予定です。

---

# client

TODO

---

# server

## download golang package

go modules を使用して golang package を download できます。

```
cd grpc-client-server-sample/go-server

go mod download
```

## run gRpc Server

```
cd grpc-client-server-sample/go-server

go run cmd/grpc_main.go
```

---

# tools

## install proto buffers

以下 URL より proto buffers をダウンロードして grpc-client-server-sample/tools/grpc 配下にリネームして配置してください。
https://github.com/protocolbuffers/protobuf/releases

## protoc-gen-go install

以下コマンドで protoc-gen-go をインストールし、バイナリを path へ設定してください。

```
go install github.com/golang/protobuf/protoc-gen-go
```

## generate proto file

```
cd grpc-client-server-sample/tools/grpc

protoc/bin/protoc --go_out=plugins=grpc:../../go-server/internal/apps/ protos/*.proto
```
