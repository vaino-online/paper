# 📜 Paper
Taking note taking to the next level.

## 🧱 Development

Paper is built using [Protocol Buffers](https://developers.google.com/protocol-buffers?authuser=1) and [gRPC](https://grpc.io). To get started, check out the [Quick Start Guide](https://grpc.io/docs/languages/go/quickstart/) for installation instructions.

After installing Go and the protocol buffer compiler, we can compile the `notes.proto` definition file with the following command: 

```
$ protoc --go_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_out= \
  --go-grpc_opt=paths=source_relative \
  notes/notes.proto
```

The compiler generates a few `.pb.go` files in the same folder as the source `.proto` file. Read more about what gets compiled  [over here](https://developers.google.com/protocol-buffers/docs/reference/go-generated?authuser=1).

After successfully generating the `.pb.go` files, we can use `notes_server/main.go` to start up a TCP gRPC server. The default port it listens on is `50051` but can be customized with the `-port` flag.

```
$ go run ./notes_server/main.go -port 50052
2022/09/19 15:49:48 Started notes server at [::]:50052
```

## 📝 Tasks

- [ ] Document the available server actions and the basic workflow

## 🤷‍♂️ Authors

🧠 Nico Hämäläinen ([@vainonico](https://twitter.com/vainonico))

## ⏳ Version History

* 0.0.1
  * Initial project setup

**References and more information:**

- [The Go programmming language](https://go.dev/)
- [gRPC: A high performance, open source universal RPC framework](https://grpc.io/)

## ©️ License

This project is licensed under the MIT License - see the LICENSE file for details.