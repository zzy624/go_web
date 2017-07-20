#!/usr/bin/env bash

all: dev run

fmt:
	goimports -l -w  ./

usevendor:
	#vendor = $(shell if [  -d "vendor" ]; then echo "exist"; else echo "notexist"; fi;)
	if [ ! -d "vendor" ]; then  mv vendor.tmp vendor ; fi;
	govendor add +e
	govendor remove +u
#	govendor update code.qschou.com/qschou/go_common/...

novendor:
	#vendor = $(shell if [  -d "vendor" ]; then echo "exist"; else echo "notexist"; fi;)
	if [  -d "vendor" ]; then  mv vendor vendor.tmp ; fi;

install: fmt clean usevendor

clean:
	rm -rf output/conf/
gotest:
	 go test ./... | grep -v "^ok "

dev: install
	go build -o output/bin/go_web ./app

test: install
	go build -tags qsc_test -o output/bin/go_web ./app

release: install
	go build -tags qsc_live -o output/bin/go_web ./app
run:
	output/bin/go_web
