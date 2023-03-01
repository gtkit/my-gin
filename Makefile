.PHONY: server gormgen linuxs lint go.mod check-diff

LocalApp=office
LinuxApp=officeaid


server: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} server  -c dev

gormgen: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} gormgen  -c dev

linuxs:tool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter  -ldflags "-s -w" -o ${LinuxApp} main.go && upx -9 ${LinuxApp}

tool:
	go vet ./...
	gofmt -l -w .

weight:
	goweight

go.mod:
	go mod tidy -v

check-diff:
	git diff --exit-code ./go.mod # check no changes
