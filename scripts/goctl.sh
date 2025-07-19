goctl api go -api *.api -dir ./  --style=go_zero

goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=go_zero

