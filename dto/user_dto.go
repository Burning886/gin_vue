// @Author:WY
// @Time:2020/4/7 16:50
package dto

import "gin_vue/models"

type Userdto struct{
	Name string `json:"name"`
	Telephone string `json:"telephone"`
}

func ToUserDto(user models.User)Userdto{
	return Userdto{
		Name:user.Name,
		Telephone: user.Telephone,
	}
}