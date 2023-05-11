.PHONY: usage  build run forever cron

PWD := $(shell pwd)
DS := /
SERVICE := palmon
COORD:=coordinator
WORKER:=worker
GCFLAG := -gcflags='all=-N -l'
PID=./logs/$(SERVICE).pid
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

default: usage

usage:
	@echo
	@echo "-> usage:"
	@echo "make build \t\t\t 同时编译 coordinator 和 worker"
	@echo "make vet  \t\t\t 当前项目代码静态检查"
	@echo "make coordinator \t\t\t 构建coordinator"
	@echo "make worker  \t\t 构建 worker"
	@echo "make clean \t\t\t 清理"
env:
	@go mod tidy
vet:
	@echo '->[$(SERVICE)] 正在检查代码'
	@go vet ./
	@echo '->[$(SERVICE)] 检查代码完成'

coordinator: vet
	@echo '->[$(SERVICE).$(COORD)] 正在构建'
	@$(if $(wildcard bin), , mkdir -p bin)
	@go build -o bin$(DS)$(COORD) $(GCFLAG) cmd/coordinator.go
	@echo '->[$(SERVICE).$(COORD)] 构建完成'
worker: vet
	@echo '->[$(SERVICE).$(WORKER)] 正在构建'
	@$(if $(wildcard bin), , mkdir -p bin)
	@go build -o bin$(DS)$(WORKER) $(GCFLAG)  cmd/worker.go
	@echo '->[$(SERVICE).$(WORKER)] 构建完成'
build: worker coordinator
	@echo '->[$(SERVICE)] 构建完成'
clean:
	@rm -rf  $(PWD)$(DS)bin








