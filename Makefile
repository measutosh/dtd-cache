# .PHONY: run
# run: main
#	 ./$<
#
# main: *.go go.mod
# 	go build -o $@ .
#
# .PHONY: all
# all: main

build:
	go build -o bin/dtb-cache

run: build
	./bin/dtb-cache