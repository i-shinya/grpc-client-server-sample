# grpc-client-server-sample
golangでgrpcサンプルを実装するリポジトリ
クライアントとサーバの実装サンプルを実装していく予定です。

--- 

# client
TODO

---

# server
TODO

---

# tools

## install proto buffers

以下URLよりproto buffersをダウンロードしてgrpc-client-server-sample/tools/grpc配下にリネームして配置してください。
https://github.com/protocolbuffers/protobuf/releases

## protoc-gen-go install 

```
go install github.com/golang/protobuf/protoc-gen-go
```

## generate proto file

```
cd grpc-client-server-sample/tools/grpc

protoc/bin/protoc --go_out=plugins=grpc:./ protos/*.proto
```

