all: build_fwserver build_fwdaemon

build_fwserver:
	cd ${GOPATH} && go get -v ...fwserver && go build ...fwserver

build_fwdaemon:
	cd ${GOPATH} && go get -v ...fwdaemon && go build ...fwdaemon

adduser:
	psql -U fwdb -W fwdb -h localhost
	INSERT INTO users (api_key,updated_at,last_checked_at) VALUES ('5c', NOW(),NOW());

update:
	curl -d '{"api_key": "5c"}' http://fnord.duckdns.org:8000/update

