CURRENT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
VERSION := $(shell git describe --tag --always --long)

build: ssh-crypt ssh-crypt.exe ssh-crypt_darwin

ssh-crypt: $(wildcard *.go) $(wildcard **/*.go)
	go get ./...
	go get github.com/stretchr/testify/assert
	go get github.com/mattn/goveralls
	go test -v github.com/suside/ssh-crypt/lib -coverprofile=main.coverprofile
	go build -o ssh-crypt -i -ldflags "-s -w -X main.version=${VERSION}"
	${GOPATH}/bin/goveralls -coverprofile=main.coverprofile -service travis-ci || true

ssh-crypt.exe: ssh-crypt
	env GOOS=windows go build -o ssh-crypt.exe -i -ldflags "-s -w -X main.version=${VERSION}"

ssh-crypt_darwin: ssh-crypt
	env GOOS=darwin go build -o ssh-crypt_darwin -i -ldflags "-s -w -X main.version=${VERSION}"
