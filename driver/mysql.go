package driver


import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(){
	dsn := "root:@tcp(127.0.0.1:3306)/vac?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err=gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err.Error())
	}
	initMigrate()
}
func initMigrate(){
	models:=registerEntities()
	for _,model:=range models{
		DB.AutoMigrate(&model.Model)
	}
}