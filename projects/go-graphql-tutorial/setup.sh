#!/bin/sh

go install github.com/99designs/gqlgen@latest
gqlgen init
gqlgen generate

docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=dbpass -e MYSQL_DATABASE=hackernews -d mysql:latest
docker exec -it mysql bash
mysql -u root -p
CREATE DATABASE hackernews

go get -u github.com/go-sql-driver/mysql
go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o "$GOPATH"/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
cd internal/pkg/db/migrations/ || exit
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table

migrate -database mysql://root:dbpass@/hackernews -path internal/pkg/db/migrations/mysql up
