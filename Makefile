BUF_VERSION:=1.3.1
SVC_PROTO_PATH:=api/proto/example/foo

generate:
	podman run -v $$(pwd):/src:z -w /src --rm docker.io/bufbuild/buf:$(BUF_VERSION) --path $(SVC_PROTO_PATH) generate