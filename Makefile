COMMIT			:=		$(shell git rev-parse --short HEAD)
VERSION			:=		$(shell git describe --tags --always)
REGISTRY		:=		233121219514.dkr.ecr.us-east-1.amazonaws.com
IMAGE_NAME		:=		${REGISTRY}/numeral

image:
	docker build --build-arg commit=${COMMIT} --build-arg version=${VERSION} -t ${IMAGE_NAME} .

build: 
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build -ldflags="-w -s -X main.commit=${COMMIT} -X main.version=${VERSION}" -o ./bin/main ./cmd/container

.PHONY: image build