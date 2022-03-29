# go-docker-graphql-sample-2
golang docker GraphQL サンプル


### migration
- migrate

```
migrate -source file://database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' up
```

- reset

```
migrate -source file://database/migrations -database 'mysql://user:pass@tcp(127.0.0.1:3306)/20220328_GO_GRAPHQL_DB' down
```


- 参考
  - https://github.com/golang-migrate/migrate
  - https://ichi-station.com/golang-migrate/