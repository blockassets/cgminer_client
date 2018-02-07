DATE=$(shell date -u '+%Y-%m-%d %H:%M:%S')
COMMIT=$(shell git log --format=%h -1)
VERSION=client.version=${TRAVIS_BUILD_NUMBER} ${COMMIT} ${DATE}
COMPILE_FLAGS=-ldflags="-X '${VERSION}'"

build:
	@go build ${COMPILE_FLAGS}

test:
	@go test .

dep:
	@dep ensure

all: clean test build
