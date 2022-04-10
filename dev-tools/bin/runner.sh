#!/bin/bash

#DB_PASSWORD="pass"
#TEST_DB_ENV_STRING="test"
# コマンド実行時の第2引数以下をARGに格納
#ARG_SECOND=${@:2}
#ARG_THIRD=${@:3}
#ARG_FORTH=${@:4}

case ${1} in
# テーブルからentityを自動生成 (事前にdbコンテナを立ち上げておくこと)
"create:entity")
 sqlboiler mysql -c app/database.toml -o app/database/entity -p entity --no-tests --wipe
 ;;
"migrate:up")
 migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' up
 ;;
"migrate:down")
 migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' down
 ;;
"lint")
 # shellcheck disable=SC2164
 cd app
 staticcheck ./...
 ;;
"test:all")
 # shellcheck disable=SC2164
 cd app
 go test -v ./...
 ;;
#"test")
# # shellcheck disable=SC2164
# cd app
# go test -v "${ARG_SECOND}" "${ARG_THIRD}" "${ARG_FORTH}"
# ;;
"gql")
# shellcheck disable=SC2164
 cd app
 gqlgen generate
 ;;
esac