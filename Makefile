GO_VERSION = 1.17

tidy:
	go mod tidy

gen: tidy
	buf generate && buf lint