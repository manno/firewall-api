export GOPATH := $(shell pwd )
export GOBIN := $(shell pwd )/bin

all: fwserver fwdaemon

fwserver:
	cd src/mm/fwserver/ && go get
	go build -a mm/fwserver

fwdaemon:
	cd src/mm/fwdaemon && go get
	go build -a mm/fwdaemon


adduser:
	psql -U fwdb -W fwdb -h localhost
	INSERT INTO users (api_key,updated_at,last_checked_at) VALUES ('5c', NOW(),NOW());

update:
	curl -d '{"api_key": "5c"}' http://fnord.duckdns.org:8000/update

