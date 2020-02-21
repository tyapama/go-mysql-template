# !/bin/bash

apt-get update -qq && \
apt-get install -y mysql-client vim && \
mkdir /go/src /go/pkg /go/bin && \
go mod download

# MySQLサーバーが起動するまでmain.goを実行せずにループで待機する
until mysqladmin ping -h db -P 3306 --silent; do
  echo 'waiting for mysqld to be connectable...'
  sleep 2
done

echo "app is starting...!"
exec go run main.go