// @Author:WY
// @Time:2020/4/7 16:31
package middleware

import (
	"gin_vue/dao"
	"gin_vue/models"
	"gin_vue/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取authorization header
		tokenString := c.GetHeader("authorization")
		//验证格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//通过验证获取claims中的userId
		userId:=claims.UserId
		var user models.User
		dao.DB.First(&user,userId)
		//用户不存在
		if user.ID<0{
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			c.Abort()
			return
		}
		//用户存在,将用户信息写入上下文
		c.Set("user",user)
		c.Next()

	}
}


