//model.go
package blog

type Context struct {
	Config
	User
	Db
}

/**
 * 配置文件中的内容
 */
type Config struct {
	Port int
}

type User struct {
	Id       int
	Name     string
	Password string
	Email    string
}

type Article struct {
	Id        int
	Title     string
	Content   string
	WriteTime string
}

/**
 * 模型层通用接口
 */
type Model interface {
	Insert()
	Update()
	Delete()
	Read()
}
