REPOSITORY_TAG := $(shell git symbolic-ref --short HEAD)
PROJECT_DIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
CODE_GENERATOR_VERSION := "v0.30.1"
PROTOC_GEN_GOGO_VERSION := "v1.3.2"

vendor:
	rm -r "${PROJECT_DIR}/vendor/" || true
	go mod tidy
	go mod vendor

code-generator:
	if [ -d "${PROJECT_DIR}/../code-generator" ]; then rm -rf "${PROJECT_DIR}/../code-generator"; fi
	git clone https://github.com/kubernetes/code-generator.git "${PROJECT_DIR}/../code-generator"
	git -C "${PROJECT_DIR}/../code-generator" checkout ${CODE_GENERATOR_VERSION}
	cd "${PROJECT_DIR}/../code-generator" && go mod tidy
	cd "${PROJECT_DIR}/../code-generator" && go build -o ${PROJECT_DIR}/bin/go-to-protobuf "./cmd/go-to-protobuf/main.go"

protoc-gen-gogo:
	if [ -d "${PROJECT_DIR}/../protobuf-gogo" ]; then rm -rf "${PROJECT_DIR}/../protobuf-gogo"; fi
	git clone https://github.com/gogo/protobuf.git "${PROJECT_DIR}/../protobuf-gogo"
	git -C "${PROJECT_DIR}/../protobuf-gogo" checkout ${PROTOC_GEN_GOGO_VERSION}
	cd "${PROJECT_DIR}/../protobuf-gogo" && go mod tidy
	cd "${PROJECT_DIR}/../protobuf-gogo" && go build -o ${PROJECT_DIR}/bin/protoc-gen-gogo "./protoc-gen-gogo/main.go"

apiserver-build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/apiserver-linux-amd64 ./cmd/apiserver/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o ./bin/apiserver-linux-arm64 ./cmd/apiserver/main.go

lifecycle-build:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./bin/lifecycle-linux-amd64 ./cmd/lifecycle/main.go
	CGO_ENABLED=0 GOARCH=arm64 GOOS=linux go build -o ./bin/lifecycle-linux-arm64 ./cmd/lifecycle/main.go

apiserver-buildah:
	buildah manifest rm apiserver 2> /dev/null || true
	buildah manifest create apiserver
	buildah bud --tag "registry.scaleship.io/kubeception/apiserver:${REPOSITORY_TAG}" --manifest apiserver --arch amd64 --os linux --build-arg BINARY_OS=linux --build-arg BINARY_ARCH=amd64 .
	buildah bud --tag "registry.scaleship.io/kubeception/apiserver:${REPOSITORY_TAG}" --manifest apiserver --arch arm64 --os linux --build-arg BINARY_OS=linux --build-arg BINARY_ARCH=arm64 .
	buildah manifest push --all apiserver "docker://registry.scaleship.io/kubeception/apiserver:${REPOSITORY_TAG}"

lifecycle-buildah:
	buildah manifest rm lifecycle 2> /dev/null || true
	buildah manifest create lifecycle
	buildah bud --tag "registry.scaleship.io/kubeception/lifecycle:${REPOSITORY_TAG}" --manifest lifecycle --arch amd64 --os linux --build-arg BINARY_OS=linux --build-arg BINARY_ARCH=amd64 .
	buildah bud --tag "registry.scaleship.io/kubeception/lifecycle:${REPOSITORY_TAG}" --manifest lifecycle --arch arm64 --os linux --build-arg BINARY_OS=linux --build-arg BINARY_ARCH=arm64 .
	buildah manifest push --all lifecycle "docker://registry.scaleship.io/kubeception/lifecycle:${REPOSITORY_TAG}"

clean:
	go clean
	cd ./bin/ && find . -maxdepth 1 -type f -name 'apiserver-*' -delete
	cd ./bin/ && find . -maxdepth 1 -type f -name 'lifecycle-*' -delete
	buildah manifest rm apiserver 2> /dev/null || true
	buildah manifest rm lifecycle 2> /dev/null || true