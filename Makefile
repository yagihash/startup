export GO111MODULE := on

.PHONY: download
download:
	@ go mod download

.PHONY: run
run:
	@ go run main.go

