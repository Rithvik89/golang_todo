package authmanager

import (
	"golang_start/entities"
	kvstore "golang_start/pkg/kv_store"

	"gorm.io/gorm"
)

type AuthManager struct {
	// redis kv store  ....
	kv *kvstore.RedisKVStore
	// GORM variable ....
	db *gorm.DB
}

func New(kv *kvstore.RedisKVStore, db *gorm.DB) *AuthManager {
	return &AuthManager{
		kv: kv,
		db: db,
	}
}

func (am *AuthManager) RegisterUser(mailId string, password string) (string, error) {
	var Session string
	DB := am.db.Create(&entities.Users{
		MailId:   mailId,
		Password: password,
	})
	if DB.Error != nil {
		return Session, DB.Error
	}
	return Session, nil
}

func (am *AuthManager) LoginUser(mailId string, password string) {

}

func (am *AuthManager) LogoutUser() {

}
