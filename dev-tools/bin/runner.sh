#!/bin/bash

# シェルスクリプト サンプル
#ARD=${@:2}

#DB_PASSWORD="pass"
TEST_DB_ENV_STRING="test"

case ${1} in
"test:all")
#  cd app
#  go test -v ./...
  docker-compose up -d server db-test
  docker-compose exec -e TEST_DB_ENV="${TEST_DB_ENV_STRING}" server go test -v ./...
#  go test -v ./...
#  docker-compose exec -e server go test -v ./...
  docker-compose stop
  ;;
esac