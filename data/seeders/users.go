package seeders

import (
	"github.com/jinzhu/gorm"
	"go-oauth/data/models/users"
	"go-oauth/helpers/crypto"
)

func CreateUser(db *gorm.DB, firstName string, lastName string, username string, password string, role string) error {
	return db.Create(&users.User{
		FirstName: firstName,
		LastName: lastName,
		Username: username,
		Password:      crypto.HashAndSalt([]byte(password)),
		Role:      users.UserRole{RoleName: role},
	}).Error
}