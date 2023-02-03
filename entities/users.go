package entities

import "time"

type Users struct {
	Id       uint `gorm:"primary_key;auto_increment"`
	MailId   string
	Password string
	CreateAt time.Time
}
