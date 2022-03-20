package repoMySqlImpl

import (
	"database/sql"
	"student_score_sword/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var sqlDB *sql.DB

type Database struct {
	MaxConn int
	MaxOpen int
}

var DatabaseConfig = new(Database) //设置全局的引用型指针变量

func GetConn() (*gorm.DB, error) {
	dsn := "root:root@tcp(192.168.31.106:32769)/student_score_sword?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db, error
}

func init() {
	conn, _ := GetConn()
	conn.AutoMigrate(&models.Class{})
	conn.AutoMigrate(&models.Exam{})
	conn.AutoMigrate(&models.Score{})
	conn.AutoMigrate(&models.Student{})
	conn.AutoMigrate(&models.Teacher{})
}
