// @Author:WY
// @Time:2020/4/7 15:29
package controller

//viewfunction
import (
	"fmt"
	"gin_vue/dao"
	"gin_vue/dto"
	"gin_vue/models"
	"gin_vue/response"
	"gin_vue/utils"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	//获取参数
	var requestUser = models.User{}
	c.Bind(&requestUser)
	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(name)
	fmt.Println(telephone)
	fmt.Println(password)
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 || len(password) > 12 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度在6-12之间")
		return
	}
	//如果名称没有传,给一个10位额随机字符串
	if len(name) == 0 {
		name = utils.RandomString(10)
	}
	log.Print(name, telephone, password)
	//判断手机号是否存在
	if isTelephoneExist(dao.DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, gin.H{"code": 422, "msg": "用户已存在"}, "")
		return
	}
	//创建用户
	//密码加密
	hasePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 500, "msg": "加密错误"})
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasePassword),
	}
	fmt.Printf("%v\n", newUser)
	err = dao.DB.Create(&newUser).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// else {
	//	c.JSON(http.StatusOK, newUser)
	//}
	//发放token
	token, err := utils.ReleaseToken(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	response.Success(c, gin.H{"token": token}, "注册成功")
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Login(c *gin.Context) {
	//获取参数
	var requestUser = models.User{}
	c.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password
	fmt.Println(telephone)
	fmt.Println(password)
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 || len(password) > 12 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码长度在6-12之间")
		return
	}
	//判断手机号是否存在
	var user models.User
	dao.DB.Where("telephone=?", telephone).First(&user)
	if user.ID < 0 {

		c.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}
	//判断密码是否正确1. 原始密码,2.用户传过来的密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
	}
	//发放token
	token, err := utils.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error:%v", err)
		return
	}
	//返回结果
	//c.JSON(http.StatusOK,gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	//	"msg":"登录成功",
	//})
	response.Success(c, gin.H{"token": token}, "登录成功")
}
func Info(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(models.User))},
	})
}
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
