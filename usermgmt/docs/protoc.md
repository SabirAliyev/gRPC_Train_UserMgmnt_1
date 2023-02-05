1. After creation main directories run command to initiate module:

    `go mod init example.com/go-usermgmt-grpc`
 

2. After creation `.proto` file run this command from project main directory to generate go files by usermgmt.proto file:

   `protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative usermgmt/usermgmt.proto`


3. After this we will get two files: usermgmt.pb.go, usermgmt_grpc.pb.go.

The `usermgmt.pb.go` file provides logic for serialising and de-serialising the messages? that we defined in service definition.
The `usermgmt_grpc.go` file includes generated client and server code, that we need to implement in our own client and server programs.

4. After creation all files run this command to download all packages:

   `go mod tidy`

