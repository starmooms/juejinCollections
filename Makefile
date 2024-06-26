
APP = juejin_collections
## MKFILE_FILE := $(abspath $(lastword $(MAKEFILE_LIST)))
## MKFILE_DIR := $(dir $(MKFILE_FILE))
## MKFILE_DIR := $(patsubst %/,%, $(dir $(MKFILE_FILE)))
ROOT_PATH := $(CURDIR)
FRONTEND_PATH := ./frontend

## 普通build
build:
	@go build -o ${APP}

## go编译环境变量
## 官方文档 https://go.dev/doc/install/source#environment
# CGO_ENABLED 

# GOOS 目标可执行程序运行操作系统
# 支持 darwin，freebsd，linux，windows

# GOARCH 目标可执行程序操作系统构架
# 包括 386，amd64，arm

## https://github.com/mattn/go-sqlite3
## go-sqlite3 可能需要cgo编译，需要对应系统的c编译器，所有CGO_ENABLED=0设置为0，但文件可能无法运行。。。

## linux: 编译打包linux
go-linux:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-linux64 ./main.go
 
## win: 编译打包win
go-win:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-win64.exe -ldflags="-H=windowsgui"
 
## mac: 编译打包mac
go-mac:
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build $(RACE) -o ./bin/${APP}-darwin64 ./main.go
 
## 编译win，linux，mac平台
go-all: win linux mac

## 检查go模块 进行安装或删除
tidy:
	@go mod tidy

## 清理二进制文件
clean:
	@if [ -f ./bin/${APP}-linux64 ] ; then rm ./bin/${APP}-linux64; fi
	@if [ -f ./bin/${APP}-win64.exe ] ; then rm ./bin/${APP}-win64.exe; fi
	@if [ -f ./bin/${APP}-darwin64 ] ; then rm ./bin/${APP}-darwin64; fi

help:
	@echo "make mac - 编译 Go 代码, 生成mac的二进制文件"
	@echo "make win - 编译 Go 代码, 生成window的二进制文件"

## frontend编译
front-build:
	cd ./frontend && yarn build

## 用statik 打包静态文件
static-build:
	@if [ -f ./statikFs/temp/ ] ; then rm ./statikFs/temp/; fi
	mkdir -p ./statikFs/temp/frontend/ ./statikFs/temp/collectReq/
	cp ./frontend/dist/ ./statikFs/temp/frontend/ -r
	cp ./collectReq/mock.json ./statikFs/temp/collectReq/mock.json -r
	statik -src=${ROOT_PATH}/statikFs/temp -dest=${ROOT_PATH}/statikFs -f
	rm -rf ./statikFs/temp/

pre-build: front-build static-build

win: pre-build go-win
mac: pre-build go-mac
linux: pre-build go-linux
all: pre-build go-all

dev:
	cd FRONTEND_PATH && yarn dev

dev-server:
	go run .

dev-server-fresh: 
	fresh -c fresh.conf

# 安装依赖
dep-download:
	go mod download && cd FRONTEND_PATH && yarn install


.PHONY: \
	go-win \
	go-linux \
	go-mac \
	go-all \
	tidy \
	clean \
	help \
	front-build \
	static-build \
	pre-build \
	win \
	mac \
	linux \
	all \
	dev \
	dev-server \
	dev-server-fresh \
	dep-download \


# // http://www.45fan.com/article.php?aid=1D7T0Iy4Q43XhrJH
# APP = task

# build: 
# 	go build -o juejin_collections.exe