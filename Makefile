makefile_dir		:= $(abspath $(shell pwd))

.PHONY: list bootstrap init build

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	rm -rf objecttype/*.enumer.go
	rm -rf objecttype/*.optioner.go
	rm -rf stringformat/*.enumer.go
	rm -rf stringformat/*.optioner.go

build:
	go build

install:
	go install

generate:
	go generate ./...