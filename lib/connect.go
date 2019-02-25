package lib

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const driver = "mysql"

func joinDataSourceName(conn map[string]interface{}) string {
	username := conn["username"].(string) // 这个语法以前没接触过
	password := conn["password"]
	host := conn["host"]
	port := fmt.Sprintf("%v", conn["port"]) // 值为 interface{} 类型，所以不能用 strconv 包来转换
	database := conn["database"]

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, database)
}

func Connect(conn *map[string]interface{}) (*sqlx.DB, error) {

	dataSourceName := joinDataSourceName(*conn)
	// root:123456@tcp(localhost:3306)/test
	fmt.Println(dataSourceName)
	// 内部会执行Open()和Ping()，Ping()失败时会关闭连接
	db, err := sqlx.Connect(driver, dataSourceName)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Query(db *sqlx.DB, sql string, params map[string]interface{}) (*sqlx.Rows, error) {
	// 执行SQL查询，返回 sqlx.Rows 对象指针
	return db.NamedQuery(sql, params)
}
