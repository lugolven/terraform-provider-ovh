export ROOT_DIR:=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
export GOBIN=${ROOT_DIR}/bin

go-files: **.go go.*

${GOBIN}/terraform-provider-ovh: go-files
	go install


build: ${GOBIN}/terraform-provider-ovh

tests: go-files
	TF_ACC=True go test -v ./...

lint:
	golangci-lint run --config ${ROOT_DIR}/.golangci.yml

ci: tests lint build
