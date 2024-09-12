.PHONY: proto help

PROTO_DIR=.

proto:
	protoc \
	-I=proto/$(PROTO_DIR) \
	--go_out=pb/$(PROTO_DIR) \
	--go_opt=paths=source_relative \
	$(shell find $(PROTO_DIR)  -iname "*.proto" -not -path "./google/*")

help:
	@echo "make proto - 生成pb代码" $(shell find $(PROTO_DIR) -iname "*.proto")

google:
	$(shell cd $(PROTO_DIR)"/proto" && git submodule add -f git@github.com:xiaozhaofu/google.git google )
