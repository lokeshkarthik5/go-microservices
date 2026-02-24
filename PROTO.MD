protoc \
  --go_out=. \
  --go-grpc_out=. \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  proto/user.proto

Run this to generate proto.go files

Replace with user.proto with your file