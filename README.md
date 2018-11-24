# pb_walker
Parses Protocol Buffer files and generates codes for Go and Java. It produce RPC service Interfaces, default implementations, idiomatic code. The generated code is like gRPC but with more capabilities and customized network and error handling. Network has been implemented on top of websocket and on fly can changes to HTTP. The result is a consistent code base, that reduce a lot of boilerplate codes, and has significantly reduced API implementation time, and freedom to test with more features, and more API methods. If API method is not implemented it will be an compile time error. Single source of trust.

### Sample output
Real world out put code of this code genrator:
From this `.proto` files: https://github.com/jozn/sun2/tree/master/shared/proto

Android: https://github.com/jozn/ms_native/tree/master/app/src/main/java/ir/ms/pb

Fo backend Go: 
- https://github.com/jozn/sun2/blob/master/shared/x/RPC_HANDLERS.java
- https://github.com/jozn/sun2/blob/master/shared/x/pb__gen_ant.go
- https://github.com/jozn/sun2/blob/master/shared/x/pb__gen_ant_empty.go
- https://github.com/jozn/sun2/blob/master/shared/x/pb__gen_enum.proto
- https://github.com/jozn/sun2/blob/master/shared/x/rpc_client.go
- https://github.com/jozn/sun2/blob/master/shared/x/flat.go

