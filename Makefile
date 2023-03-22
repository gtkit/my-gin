.PHONY: server gormgen linuxs lint go.mod check-diff

LocalApp=office
LinuxApp=officekey


server: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} server  -c dev

gormgen: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} gormgen  -c dev

rsa: tool
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} newrsa  -c dev -a d

linuxs:tool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter  -ldflags "-s -w" -o ${LinuxApp} main.go && upx -9 ${LinuxApp}

tool:
	go vet ./...
	gofmt -l -w .

weight:
	goweight

go.mod:
	go mod tidy -v

go.update:
	go get -u
	go mod tidy -v

check-diff:
	git diff --exit-code ./go.mod # check no changes
