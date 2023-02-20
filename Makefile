.PHONY: server

APP=office


server: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${APP} main.go; ./${APP} server  -c dev
tool:
	go vet ./...
	gofmt -l -w .
