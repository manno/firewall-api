export GOPATH := $(shell pwd )
export GOBIN := $(shell pwd )/bin

all: fwserver fwdaemon

fwserver:
	cd src/mm/fwserver/ && go get
	go build -a mm/fwserver

fwdaemon:
	cd src/mm/fwdaemon && go get
	go build -a mm/fwdaemon
