package service

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

var (
	resource  *dockertest.Resource
	pool      *dockertest.Pool
	con       *sql.DB
	dbTestErr error
)

func TestMain(m *testing.M) {
	//	beforeAll
	fmt.Println("beforeAll")
	beforeAll()
	m.Run()
	//	afterAll
	fmt.Println("afterAll")
	afterAll()
}

func createContainer() {
	// testutil.goの絶対パスを取得
	_, fileName, _, _ := runtime.Caller(0)

	// Dockerとの接続
	pool, dbTestErr = dockertest.NewPool("")
	pool.MaxWait = time.Minute * 2
	if dbTestErr != nil {
		log.Fatalf("Could not connect to docker: %s", dbTestErr)
	}

	// Dockerコンテナ起動時の細かいオプションを指定する
	runOptions := &dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8.0",
		Env: []string{
			"MYSQL_ROOT_PASSWORD=secret",
		},
		Mounts: []string{
			fmt.Sprintf("%s/../../mysql/mysql.cnf:/etc/mysql/conf.d/mysql.cnf", filepath.Dir(fileName)),
			fmt.Sprintf("%s/../../mysql/db:/docker-entrypoint-initdb.d", filepath.Dir(fileName)), // コンテナ起動時に実行するSQL
		},
		Cmd: []string{
			"mysqld",
			"--character-set-server=utf8mb4",
			"--collation-server=utf8mb4_unicode_ci",
		},
	}

	// コンテナを起動
	resource, dbTestErr = pool.RunWithOptions(runOptions)
	if dbTestErr != nil {
		log.Fatalf("Could not start resource: %s", dbTestErr)
	}
}

func closeContainer() {
	//	コンテナの終了
	if dbTestErr = pool.Purge(resource); dbTestErr != nil {
		log.Fatalf("Could not purge resource: %s", dbTestErr)
	}
}

func connectDB(resource *dockertest.Resource, pool *dockertest.Pool) error {
	// DB(コンテナ)との接続
	if poolErr := pool.Retry(func() error {
		// DBコンテナが立ち上がってから疎通可能になるまで少しかかるのでちょっと待ったほうが良さそう
		time.Sleep(time.Second * 20)

		var dbErr error
		con, dbErr = sql.Open("mysql", fmt.Sprintf("root:secret@(localhost:%s)/20220328_GO_GRAPHQL_DB?charset=utf8mb4&parseTime=True&loc=Local", resource.GetPort("3306/tcp")))
		if dbErr != nil {
			return dbErr
		}
		return dbErr
	}); poolErr != nil {
		log.Fatalf("Could not connect to docker: %s", poolErr)
		return poolErr
	}
	return nil
}

func RunWithDB(t *testing.T, name string, f func(t *testing.T, db *sql.DB)) {
	beforeEach()
	// テスト実行
	t.Run(name, func(t *testing.T) {
		f(t, con)
	})
}

func execSQLScript(path string) error {
	content, fileErr := ioutil.ReadFile(path)
	if fileErr != nil {
		return fileErr
	}
	_, dbTestErr = con.Exec(bytes.NewBuffer(content).String())
	if dbTestErr != nil {
		return dbTestErr
	}
	return nil
}

func createTestData() error {
	if dbTestErr = testdata.CreateTestData(con); dbTestErr != nil {
		return dbTestErr
	}
	return nil
}

func beforeAll() {
	// コンテナ起動
	createContainer()
	// テーブル作成
	_, fileName, _, _ := runtime.Caller(0)
	dbTestErr = connectDB(resource, pool)
	if dbTestErr != nil {
		log.Fatalf("db connect error: %v", dbTestErr)
	}
	sqlFileNames := [...]string{"create_todo_table"}
	for _, sqlFileName := range sqlFileNames {
		if dbTestErr = execSQLScript(fmt.Sprintf("%s/../testdata/sql/%s.sql", filepath.Dir(fileName), sqlFileName)); dbTestErr != nil {
			log.Fatalf("%s, %v", fileName, dbTestErr)
		}
	}
}

func beforeEach() {
	_, fileName, _, _ := runtime.Caller(0)
	// データ削除
	sqlFileNames := [...]string{"truncate_todo_table"}
	for _, sqlFileName := range sqlFileNames {
		if dbTestErr = execSQLScript(fmt.Sprintf("%s/../testdata/sql/%s.sql", filepath.Dir(fileName), sqlFileName)); dbTestErr != nil {
			log.Fatalf("%s, %v", fileName, dbTestErr)
		}
	}
	// テストデータ作成
	if dbTestErr = createTestData(); dbTestErr != nil {
		return
	}
}

func afterAll() {
	// コンテナ停止
	closeContainer()
}
