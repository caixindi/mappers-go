.PHONY: build test clean docker
help:
	#  
	# Usage:
	#   make build  :  compile example code in /example/virtualdevice/
	#   make mod    :  download code dependencies.
	#   make test   :  check configmap and config,determine if it is a valid file
	#   make clean  :  clean binaries
	#   make docker :  pack example to image in /example/virtualdevice/
	@echo
build:
	go build -o ./example/virtualDevice/bin/main ./example/virtualDevice/cmd/main.go
	@echo  "[INFO] go build successful"
	@echo  "[INFO] you can cd example/virtualDevice/bin to execute it"
test:
	@cd ./internal/config/ && go test
	@ cd ./internal/configmap/ && go test
clean:
	rm -rf ./example/virtualDevice/bin/*
docker:
	docker build -f ./example/virtualDevice/Dockerfile \
		-t mapper-go-sdk-example:v0.0.1 .
