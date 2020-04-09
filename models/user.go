// @Author:WY
// @Time:2020/4/7 15:22
package models

//type User struct {
//	gorm.Model
//	Name      string `gorm:"type":varchar(20);not null`
//	Telephone string `gorm:"varchar(110);not null;unique"`
//	Password  string `gorm:"size:255;not null"`
//}
//type User struct {
//	//gorm.Model
//	ID        int `gorm:"TYPE:PRIMARY_KEY"`
//	Name      string
//	Telephone string
//	Password  string
//}
type User struct {
	ID   int    `gorm:"AUTO_INCREMENT;primary_key"`
	Name string `gorm:"TYPE: VARCHAR(255); DEFAULT:'';INDEX"`
	//todos []Todo     `gorm:"FOREIGNKEY:UserId;ASSOCIATION_FOREIGNKEY:ID"`
	//CreatedAt time.Time  `gorm:"TYPE:DATETIME"`
	//UpdatedAt time.Time  `gorm:"TYPE:DATETIME"`
	//DeletedAt *time.Time `gorm:"TYPE:DATETIME;DEFAULT:NULL"`
	Telephone string `json:"telephone"`
	Password  string `json:"password"`
}
