# go-docker-graphql-sample-2
golang docker GraphQL サンプル


### migration

- sqlファイル作成
  - テーブル作成、データ登録(Seeder)ファイルを作成する

```
migrate create -ext sql -dir マイグレーションファイルの置き場 マイグレーションファイルの名前

// 例
migrate create -ext sql -dir app/database/migrations create_todos_table
```

- migrate
  - テーブル作成、データ登録が可能
```
// dbコンテナを立ち上げる
docker-compose up -d
// migration実行
./dev-tools/bin/runner.sh migrate:up
// コンテナ停止
docker-compose down

// 基本コマンド
migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' up
```

- rollback

```
// dbコンテナを立ち上げる
docker-compose up -d
// rollback実行
./dev-tools/bin/runner.sh migrate:down
// コンテナ停止
docker-compose down

// 基本コマンド
migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' down
```


- 参考
  - https://github.com/golang-migrate/migrate
  - https://ichi-station.com/golang-migrate/

### lint
- staticcheck

```
// 便利ツール
./dev-tools/bin/runner.sh lint 

// 基本コマンド
cd app
staticcheck ./...
```

- 参考
  - https://staticcheck.io/

### test
- 全てのテスト

```
// 便利ツール
./dev-tools/bin/runner.sh test:all

// 基本コマンド
cd app
go test -v ./...
```

- 特定のファイルのテスト
```
cd app

// go test テストコードのファイル テスト対象のファイル

// 例
go test -v service/todo_service_test.go service/todo_service.go service/service_test.go
// `TestMain`はパッケージ毎に実行されるので、個別にテストする際は各パッケージのTestMainのファイルも含めること
// TestMainでテストDBコンテナを立ち上げているので、DBが関係するパッケージはこれを実装すること
```

### ORM
- model自動生成

```
// dbコンテナを立ち上げる
docker-compose up -d
// model自動生成
./dev-tools/bin/runner.sh create:entity
// コンテナ停止
docker-compose down
```

```
sqlboiler mysql -c [tomlファイルのパス] -o [成果物を置くディレクトリ名] -p [パッケージ名] --no-tests  --wip

// 例
sqlboiler mysql -c app/database.toml -o app/database/entity -p entity --no-tests --wipe
```

- 参考
  - https://zenn.dev/gami/articles/0fb2cf8b36aa09#sqlboiler

### GraphQL
- generate

```
// appディレクトリで実行
cd app

gqlgen generate
// go run github.com/99designs/gqlgen generate
```

- 参考
  - https://gqlgen.com/getting-started/
  - https://tech.layerx.co.jp/entry/2021/10/22/171242
  - https://future-architect.github.io/articles/20200609/