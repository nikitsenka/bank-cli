BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}}\
{{end}}' ./...)

BINARY_NAME=bank

bin/bank: $(BUILD_FILES)
	@go build -o "$@" ./cmd/bank

clean:
	rm -rf bin