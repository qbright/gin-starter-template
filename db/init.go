package db

import (
	"gin-starter-template/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func Init() error {
	config := config.GetConfig()
	source := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local" // DSN data source name

	user, exist := os.LookupEnv("MYSQL_USERNAME")
	if !exist {
		user = config.GetString("db.username")
	}

	pass, exist := os.LookupEnv("MYSQL_PASSWORD")
	if !exist {
		pass = config.GetString("db.pass")
	}

	addr, exist := os.LookupEnv("MYSQL_ADDRESS")

	if !exist {
		addr = config.GetString("db.addr")
	}

	dbName, exist := os.LookupEnv("MYSQL_DATABASE")

	if !exist {
		dbName = config.GetString("db.dbname")
	}

	source = fmt.Sprintf(source, user, pass, addr, dbName)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})

	if err != nil {
		log.Fatalln("DB Open error: ", err.Error())
	}

	if err != nil {
		log.Fatalln("DB AutoMigrate error: ", err.Error())
	}

	sqlDB, err := db.DB()

	if err != nil {
		log.Fatalln("DB Init error: ", err.Error())
	}
	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish init mysql with ", source)

	// user1 := modules.User{Name: "Jinzhu", Age: 18}

	// db.Create(&user1) // 通过数据的指针来创建

	// bytes, _ := json.Marshal(user1)
	// log.Println(string(bytes))

	// userJson := `{"name":"fffffxxkkk","age":20}`

	// user1 := &modules.User{}

	// json.Unmarshal([]byte(userJson), &user1)

	// db.Create(&user1) // 通过数据的指针来创建

	return nil

}

func GetDb() *gorm.DB {
	return dbInstance
}
