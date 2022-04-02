package testutil

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/YukiOnishi1129/go-docker-graphql-sample-2/app/testdata"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ory/dockertest"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

func createContainer() (*dockertest.Resource, *dockertest.Pool) {
	// testutil.goの絶対パスを取得
	_, fileName, _, _ := runtime.Caller(0)

	// Dockerとの接続
	pool, err := dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
		},
		Mounts: []string{
			fmt.Sprintf("%s/../../../mysql/mysql.cnf:/etc/mysql/conf.d/mysql.cnf", filepath.Dir(fileName)),
			fmt.Sprintf("%s/../../../mysql/db:/docker-entrypoint-initdb.d", filepath.Dir(fileName)), // コンテナ起動時に実行するSQL
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	// コンテナを起動
	resource, err := pool.RunWithOptions(runOptions)
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	return resource, pool
}

func closeContainer(resource *dockertest.Resource, pool *dockertest.Pool) {
	//	コンテナの終了
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) (*sql.DB, error) {
	// DB(コンテナ)との接続
	var db *sql.DB
	if err := pool.Retry(func() error {
		// DBコンテナが立ち上がってから疎通可能になるまで少しかかるのでちょっと待ったほうが良さそう
		time.Sleep(time.Second * 10)

		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/20220328_GO_GRAPHQL_DB?charset=utf8mb4&parseTime=True&loc=Local", resource.GetPort("3306/tcp")))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
		return nil, err
	}
	return db, nil
}

func RunWithDB(t *testing.T, name string, f func(t *testing.T, db *sql.DB)) {
	var con *sql.DB
	var err error
	recource, pool := createContainer()
	defer closeContainer(recource, pool)

	// テーブル作成
	_, fileName, _, _ := runtime.Caller(0)
	con, err = connectDB(recource, pool)
	if err != nil {
		t.Fatalf("%v", err)
	}
	sqlFileNames := [...]string{"create_todo_table", "truncate_todo_table"}
	for _, sqlFileName := range sqlFileNames {
		if err = execSQLScript(con, fmt.Sprintf("%s/../../testdata/sql/%s.sql", filepath.Dir(fileName), sqlFileName)); err != nil {
			t.Fatalf("%s, %v", fileName, err)
		}
	}

	// テストデータ作成
	if err = createTestData(con); err != nil {
		return
	}
	err = con.Close()
	if err != nil {
		return
	}

	// テスト実行
	t.Run(name, func(t *testing.T) {
		con, err = connectDB(recource, pool)
		if err != nil {
			t.Fatalf("%s, %v", fileName, err)
		}
		defer func(con *sql.DB) {
			err = con.Close()
			if err != nil {

			}
		}(con)

		f(t, con)
	})
}

func execSQLScript(con *sql.DB, path string) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = con.Exec(bytes.NewBuffer(content).String())
	if err != nil {
		return err
	}
	return nil
}

func createTestData(con *sql.DB) error {
	var err error
	if err = testdata.CreateTestData(con); err != nil {
		return err
	}
	return nil
}
