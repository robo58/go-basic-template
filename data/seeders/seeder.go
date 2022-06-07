package seeders

import (
	"github.com/jinzhu/gorm"
)

type Seeder struct {
	Name string
	Run func(*gorm.DB) error
}

func All() []Seeder {
	return []Seeder{
		Seeder{
			Name: "CreateJane",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Jane", "doe", "jdoe", "123456", "user")
				if err != nil {
					return err
				}
				return nil
			},
		},
		Seeder{
			Name: "CreateJohn",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "John", "doe", "johny", "123456", "user")
				if err != nil {
					return err
				}
				return nil
			},
		},
		Seeder{
			Name: "CreateJack",
			Run: func(db *gorm.DB) error {
				err := CreateUser(db, "Jack", "doe", "jacky", "123456", "user")
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
}