// @Author:WY
// @Time:2020/4/8 17:58
package models

import (
	"gin_vue/dao"
)

//Todo Model
//type Todo struct {
//	ID     int    `json:"id"`
//	Title  string `json:"title"`
//	Status bool   `json:"status"`
//}
type Todo struct {
	ID     int    `json:"id" gorm:"AUTO_INCREMENT"`
	Title  string `gorm:"TYPE:VARCHAR(255);DEFAULT:'';INDEX" json:"title"`
	Status bool   `json:"status"`
	UserId int    `gorm:"TYPE:int(11);NOT NULL;INDEX" json:"user_id"`
}

/*Todo增删改查*/

//创建todo
func CreateTodo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}
func GetAllTodo() (todoList []*Todo, err error) {
	if err := dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DelTodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
