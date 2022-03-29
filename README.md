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