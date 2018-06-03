gen-lib:
	protoc --plugin=${GOPATH}/bin/protoc-gen-go \
		--go_out=plugins=grpc:./src/logging/ \
		-I proto/logging \
		proto/logging/v1/*.proto
	protoc --plugin=${GOPATH}/bin/protoc-gen-go \
		--go_out=plugins=grpc:./src/logging/ \
		-I proto/logging \
		proto/logging/v2/*.proto
