.PHONY: server gormgen linuxs lint go.mod check-diff

LocalApp=mygin
LinuxApp=mygin-linux


server:
	go build -tags=jsoniter -ldflags "-s -w" -gcflags="-m"  -o ${LocalApp} main.go; ./${LocalApp} server -c dev

## gormgen 生成gorm 数据表
gormgen:
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} gormgen  -c dev

## newrsa 生成rsa公钥私钥
rsa:
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} newrsa  -c dev -a d

## 构建linux版本,需要安装upx
linuxs:tool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter  -ldflags "-s -w" -o ${LinuxApp} main.go && upx -9 ${LinuxApp}

LINT_TARGETS ?= ./...
tool: ## Lint Go code with the installed golangci-lint
	@ echo "▶️ golangci-lint run"
	golangci-lint run $(LINT_TARGETS)
	gofmt -l -w .
	@ echo "✅ golangci-lint run"

weight:
	goweight

go.mod:
	go mod tidy -v

go.update:
	go get -u
	go mod tidy -v

check-diff:
	git diff --exit-code ./go.mod # check no changes

## govulncheck 检查漏洞 go install golang.org/x/vuln/cmd/govulncheck@latest
check:
	govulncheck ./...
	gosec ./...
