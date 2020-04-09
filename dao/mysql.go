// @Author:WY
// @Time:2020/4/7 15:20
package dao

//数据库初始化
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB()(err error){
	args := "root:123456@(127.0.0.1:3306)/gin_learn?charset=utf8&parseTime=True&loc=Local"
	//args := "gin_learn:zppiiHtJFeYLRcr7@(127.0.0.1:3306)/gin_learn?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", args)//没有冒号，否则DB为空指针
	if err != nil {
		panic("连接数据库出错:" + err.Error())
	}
	return DB.DB().Ping()
}

func CloseDB() {
	DB.Close()
}
