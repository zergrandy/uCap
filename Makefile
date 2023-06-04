NAME="gbms_server"
# VERSION ?= $(shell git describe --tags --always || git rev-parse --short HEAD)
PGUSER="root"
PGPASS="root"

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(NAME) -ldflags '-X "main.releaseMode=release"' *.go
	systemctl stop gbms
	cp ./bin/$(NAME) /var/www/html/gbms/$(NAME)
	systemctl start gbms

test: 
	# Drop Database
	psql postgresql://${PGUSER}:${PGPASS}@127.0.0.1:5432/postgres -c 'DROP DATABASE IF EXISTS gbms_test WITH (FORCE);'
	# Create and Generate Test Database
	psql postgresql://${PGUSER}:${PGPASS}@127.0.0.1:5432/postgres -c 'CREATE DATABASE gbms_test;'
	psql postgresql://${PGUSER}:${PGPASS}@127.0.0.1:5432/gbms_test -f ./sql/gbms_test.sql
	go test -v