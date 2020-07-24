GOOS?=linux
GOARCH?=amd64

NAME=emitter
VERSION?=$$(git rev-parse HEAD)

REGISTRY_SERVER?=registry.videocoin.net
REGISTRY_PROJECT?=cloud

.PHONY: deploy

default: build

version:
	@echo ${VERSION}

build:
	GOOS=${GOOS} GOARCH=${GOARCH} \
		go build -mod vendor \
			-ldflags="-w -s -X main.Version=${VERSION}" \
			-o bin/${NAME} \
			./cmd/main.go

deps:
	GO111MODULE=on go mod vendor
	# https://github.com/ethereum/go-ethereum/issues/2738
	cp -r $(GOPATH)/src/github.com/ethereum/go-ethereum/crypto/secp256k1/libsecp256k1 \
	vendor/github.com/ethereum/go-ethereum/crypto/secp256k1/
	cp -r $(GOPATH)/src/github.com/karalabe/usb/hidapi \
	vendor/github.com/karalabe/usb/hidapi/
	cp -r $(GOPATH)/src/github.com/karalabe/usb/libusb \
	vendor/github.com/karalabe/usb/libusb/

lint: docker-lint

release: docker-build docker-push

docker-lint:
	docker build -f Dockerfile.lint .

docker-build:
	docker build -t ${REGISTRY_SERVER}/${REGISTRY_PROJECT}/${NAME}:${VERSION} -f Dockerfile .

docker-push:
	docker push ${REGISTRY_SERVER}/${REGISTRY_PROJECT}/${NAME}:${VERSION}

deploy:
	helm upgrade -i --wait --set image.tag="${VERSION}" -n console emitter ./deploy/helm
