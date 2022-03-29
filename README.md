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
migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' up
```

- rollback

```
migrate -source file://app/database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' down
```


- 参考
  - https://github.com/golang-migrate/migrate
  - https://ichi-station.com/golang-migrate/

### lint
- staticcheck

```
staticcheck ./...
```

- 参考
  - https://staticcheck.io/

### ORM
- model自動生成

```
sqlboiler mysql -c [tomlファイルのパス] -o [成果物を置くディレクトリ名] --no-tests

// 例
sqlboiler mysql -c app/database.toml -o models --no-tests --wipe
```

- 参考
  - https://zenn.dev/gami/articles/0fb2cf8b36aa09#sqlboiler

### GraphQL
- generate

```
// appディレクトリで実行
cd app
go run github.com/99designs/gqlgen generate
```

- 参考
  - https://gqlgen.com/getting-started/
  - https://tech.layerx.co.jp/entry/2021/10/22/171242