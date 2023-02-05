package entities

type Users struct {
	Id               uint   `gorm:"primary_key;auto_increment"`
	MailId           string `gorm:"not null"`
	Password         string
	IdentityProvider string
	SubjectId        string
}
