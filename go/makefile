install:
	@echo installing go dependencies
	@go mod download

install-tools: install
	@echo Installing tools from tools.go
	@cat ./tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go get %

# Generate interfaces
generate-proto: install-tools
	@echo Generating Proto Sources
	@mkdir -p ./gen/
	@protoc --go_out=./gen/ --validate_out="lang=go:./gen/" --go-grpc_out=./gen/ -I ../proto ../proto/*.proto
