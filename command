protoc --go_out=. person.proto

protoc \
-I=proto \
-I=$GOPATH/src \
-I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:proto \
proto/service.proto

protoc --python_out=. --mypy_out=. person.proto
python -m grpc_tools.protoc -I=. --python_out=. --grpc_python_out=. --mypy_out=. service.proto

protoc \
--plugin="protoc-gen-ts=./node_modules/ts-protoc-gen/bin/protoc-gen-ts" \
--js_out="import_style=commonjs,binary:." \
--ts_out=. \
proto/person.proto

protoc \
-I=proto \
--plugin="protoc-gen-ts=./node_modules/ts-protoc-gen/bin/protoc-gen-ts" \
--js_out="import_style=commonjs,binary:proto" \
--ts_out="service=true:proto" \
proto/service.proto

protoc \
-I=proto \
-I=$GOPATH/src \
-I=$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:proto \
proto/service.proto

protoc \
-I=proto \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--go_out=plugins=grpc:proto \
proto/service.proto

protoc \
-I=proto \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--grpc-gateway_out=logtostderr=true:proto \
proto/service.proto