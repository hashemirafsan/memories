package bootstrap

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database ...
type Database struct {}

// Connection ...
func (d *Database) Connection() (*gorm.DB) {
	db, err := gorm.Open("mysql", "docker:sph123@(172.27.0.1)/gaari?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	// defer db.Close()
	return db
}
