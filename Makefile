all: build_fwapi-frontend build_fwapi-backend

build_fwapi-frontend:
	cd ${GOPATH} && go get -v ...fwapi-frontend && go build ...fwapi-frontend

build_fwapi-backend:
	cd ${GOPATH} && go get -v ...fwapi-backend && go build ...fwapi-backend

adduser:
	psql -U fwdb -W fwdb -h localhost
	INSERT INTO users (api_key,updated_at,last_checked_at) VALUES ('5c', NOW(),NOW());

update:
	curl -d '{"api_key": "5c"}' http://fnord.duckdns.org:8000/update

