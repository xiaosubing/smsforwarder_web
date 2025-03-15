package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id       int    `gorm:"primaryKey"`
	Username string `gorm:"type:varchar(20);" json:"username"`
	Password string `gorm:"type:varchar(20);" json:"password"`
	//Phone    string `gorm:"type:varchar(20);" json:"phone"`
}

type Message struct {
	gorm.Model
	Phone     string `gorm:"size:20"`
	Sender    string `gorm:"size:128"`
	SMS       string `gorm:"type:text"`
	Code      string `gorm:"size:10"`
	SmsType   string `gorm:"size:10"`
	CreatedAt string `gorm:"size:20"`
}

func (table *User) TableName() string {
	return "user"
}

func (table *Message) TableName() string {
	return "message"
}

func GetUserInfo(username, password string) (*User, error) {
	var user User
	//sqlStr := fmt.Sprintf("select * from user where username='%s' and password='%s'", username, password)
	if err := DB.Where("username = ? AND password = ?", username, password).Last(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user, err
		} else {
			return &user, err
		}
	}
	return &user, nil
	//sqlStr := fmt.Sprintf("select * from user where username='%s' and password='%s'", username, password)
	//fmt.Println("sql 语句：", sqlStr)
	//err := DB.Where(sqlStr).First(&data).Error
	//fmt.Println(data)
	//return data, err
}
