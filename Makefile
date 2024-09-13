build-dev:
	docker build . --target development --tag openslides-projector-dev

run-tests:
	docker build . --target testing --tag openslides-projector-test
	docker run openslides-projector-test

all: gofmt gotest golinter

gotest:
	go test ./...

golinter:
	golint -set_exit_status ./...

gofmt:
	gofmt -l -s -w .
