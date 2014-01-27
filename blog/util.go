//util.go
package blog

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Db struct {
	DbName string /**数据库名称*/
	*DB           /**数据库连接*/
}

//读取配置文件，获取配置内容
