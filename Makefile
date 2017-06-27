ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: happymeter

happymeter: *.go
	go build .

clean:
	rm -f happymeter

test:
	go test .
