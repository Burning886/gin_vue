// @Author:WY
// @Time:2020/4/7 15:34
package routers

//绑定路由
import (
	"gin_vue/controller"
	"gin_vue/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.Use(middleware.CORSMiddleware())
	r.GET("/",controller.IndexHandler)
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	v1Group := r.Group("v1") //蓝图?
	{
		/*添加,http://localhost:8080/v1/todo*/
		v1Group.POST("/todo", controller.CreateTodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DelATodo)
	}
	return r
}
