export GOPATH := $(shell pwd )
export GOBIN := $(shell pwd )/bin

compile:
	cd server && go get
	go build -o bin/firewall-api-server server/*.go 
	rm bin/server
