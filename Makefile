export GOPATH := $(shell pwd )
export GOBIN := $(shell pwd )/bin

compile:
	rm -fr pkg/linux_amd64/libs 
	cd server && go get
	go build -o bin/firewall-api-server server/*.go 
	rm bin/server

cclient:
	rm -fr pkg/linux_amd64/libs 
	cd client && go get
	go build -o bin/firewall-client client/*.go 
