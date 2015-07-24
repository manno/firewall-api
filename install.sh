#!/bin/sh
host=root@104.236.54.208 

# database
ssh $host <<'EOF'
apt-get install supervisor postgresql-9.3 -y
locale-gen $LC_DATE $LANG
pg_createcluster 9.3 main --start
echo "CREATE ROLE fwapi LOGIN password 'pw12345';" | sudo -u postgres psql
echo "CREATE DATABASE fwapi ENCODING 'UTF8' OWNER fwapi;" | sudo -u postgres psql
EOF

# start app
scp $GOPATH/bin/fwapi-backend $GOPATH/bin/fwapi-frontend $host:/usr/local/bin
scp supervisord_fwapi.conf $host:/etc/supervisor/conf.d/fwapi.conf
ssh $host service supervisor restart

ssh $host <<'EOF'
apt-get install polipo
echo "INSERT INTO users (api_key,updated_at,last_checked_at) VALUES ('5c', NOW(),NOW());" | sudo -u postgres psql fwapi
echo #
echo curl -d \'{"api_key": "5c"}\' http://`echo \$SSH_CONNECTION | cut -d" " -f 3`:8000/update
EOF

