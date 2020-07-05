package infrastructure

import "github.com/jinzhu/gorm"

var db *gorm.DB

func InitDB(config string) error {
	var err error
	db, err = gorm.Open("mysql", config)
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}
