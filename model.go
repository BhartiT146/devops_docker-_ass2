package model
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type User struct {
	gorm.Model
	Id    string    `json:"id"`
	name   string `gorm:"unique" json:"name"`
	age    int    `json:"age"`
	dept   string `json:"dept"`
	subject   string `json:"subject"`
}
// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{})
	return db
}
