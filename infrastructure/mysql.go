package infrastructure

import "github.com/jinzhu/gorm"

var db *gorm.DB

// InitDB .
func InitDB(config string) error {
	var err error
	db, err = gorm.Open("mysql", config)
	if err != nil {
		return err
	}
	return nil
}

// GetDB .
func GetDB() *gorm.DB {
	return db
}
