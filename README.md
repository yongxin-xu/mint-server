## Mint Server
A lightweight server written in Golang. It is used in Tencent new graduate MiniGame Tian Wen.

Mint server is inspired by [Zinx](https://github.com/aceld/zinx), thanks so much for the zinx project.

Besides multi-thread (actually goroutines), the mint server also has the feature of read-write goroutine separation and task queue workload limitation.

It uses [Protobuf](https://github.com/golang/protobuf) to encode messages and has functionalities of signing up, signing in, and game progress saving. All data is stored in [MySQL](https://github.com/mysql/mysql-server). Users can apply a main-standby configuration on top of that. 

## Quick Start
Configuration example is as [config-example.json](https://github.com/yongxin-xu/mint-server/blob/master/config/config-example.json).

Client code example: [client.go](https://github.com/yongxin-xu/mint-server/blob/master/example/client/client.go)

Server code example: [server.go](https://github.com/yongxin-xu/mint-server/blob/master/example/server/server.go)
