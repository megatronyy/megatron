package sqlcommon

import (
	"github.com/go-xorm/xorm"
)

const Conn = "root:QU90()op@tcp(192.168.0.106:3306)/pholcus?charset=utf8"

type Config struct {
	Engine *xorm.Engine
}

//NewConfig 初始化数据库连接
func NewConfig() *Config {
	engine, err := xorm.NewEngine("mysql", Conn)
	if err != nil {
		panic(err)
	}

	return &Config{
		Engine: engine,
	}
}
