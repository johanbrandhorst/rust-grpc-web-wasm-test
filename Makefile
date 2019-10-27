generate:
	go install github.com/golang/protobuf/protoc-gen-go
	cd proto && protoc -I. library/book_service.proto --go_out=plugins=grpc,paths=source_relative:./gen/go
	# run rust wasm generator ???

generate_cert:
	cd insecure && go run "$$(go env GOROOT)/src/crypto/tls/generate_cert.go" \
		--host=localhost,127.0.0.1 \
		--ecdsa-curve=P256 \
		--ca=true
