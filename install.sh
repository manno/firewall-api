scp $GOPATH/bin/fwdaemon $GOPATH/bin/fwserver $HOST:/usr/local/bin
scp supervisord_fwdaemon.conf $HOST:/etc/supervisor/conf.d/fwdaemon
