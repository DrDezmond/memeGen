.PHONY: run
runGenerator:
							go run -v ./generatorService
.PHONY: build
build:
							go build ./cmd/apiserver