.PHONY: server gormgen linuxs lint go.mod check-diff

MacApp=mygin
LinuxApp=mygin-linux
WinApp=mygin.exe

## 构建mac版本,需要安装upx
manual:
	## 手动构建, 使用不同的json库
	#go build -tags="sonic,avx,darwin,amd64" -ldflags "-s -w" -gcflags="-m"  -o ${LocalApp} main.go; ./${LocalApp} manual -c dev
	#go build -tags=jsoniter -ldflags "-s -w" -gcflags="-m"  -o ${LocalApp} main.go; ./${LocalApp} manual -c dev
	#go build -tags=go_json -ldflags "-s -w" -gcflags="-m"  -o ${LocalApp} main.go; ./${LocalApp} manual -c dev
	go build  -ldflags "-s -w" -gcflags="-m"  -o ${MacApp} main.go; ./${MacApp} manual -c dev
server:
	go build -tags=jsoniter -ldflags "-s -w" -gcflags="-m"  -o ${MacApp} main.go; ./${MacApp} server -c dev

## 构建windows版本,需要安装upx
win:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -tags=jsoniter -ldflags "-s -w" -o ${WinApp} main.go && upx -9 ${WinApp}

## 构建linux版本,需要安装upx
linuxs:tool
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter  -ldflags "-s -w" -o ${LinuxApp} main.go && upx -9 ${LinuxApp}

## gormgen 生成gorm 数据表
gormgen:
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} gormgen  -c dev

## newrsa 生成rsa公钥私钥
rsa:
	go build -tags=jsoniter -ldflags "-s -w" -o ${LocalApp} main.go; ./${LocalApp} newrsa  -c dev -a d

LINT_TARGETS ?= ./...
tool: ## Lint Go code with the installed golangci-lint
	@ echo "▶️ golangci-lint run"
	golangci-lint run $(LINT_TARGETS)
	@ echo "✅ golangci-lint run"

weight:
	goweight
go.update:
	go get -u
	go mod tidy -v

check-diff:
	git diff --exit-code ./go.mod # check no changes

## govulncheck 检查漏洞 `go install golang.org/x/vuln/cmd/govulncheck@latest`
## gosec 检查安全漏洞 `go install github.com/securego/gosec/v2/cmd/gosec@latest`
check:
	govulncheck ./...
	gosec ./...
## betteralign 优化结构体字段排序和内存布局 `go install github.com/dkorunic/betteralign/cmd/betteralign@latest`
better:
	betteralign ./...
	betteralign -apply ./...
## gofumpt 格式化代码 `go install mvdan.cc/gofumpt@latest`
fmt:
	gofumpt -l -w .
	go mod tidy
