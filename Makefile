makefile_dir	:= $(abspath $(shell pwd))
m				?= "updates"

.PHONY: list purge build install generate test commit tag publish

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

purge:
	find . -name '*.enum.go' -type f -delete

build: generate
	go build

install: generate
	go install

generate:
	go generate ./...

test: generate
	go test ./...

push:
	git add . || true
	git commit -m "$(m)" || true
	git push origin master

tag:
	git tag -a $(tag) -m "$(tag)"
	git push origin $(tag)

tag-list:
	git fetch --tags
	git tag -l | sort -V

publish: test
	@if ack replace go.mod ;then echo 'Remove the "replace" line from the go.mod file'; exit 1; fi
	make push m=$(m)
	make tag tag=$(m)