package dao

import (
	"Todolist/config"
	"Todolist/repository/db/model"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

func MySQLInit() {
	mysqlCfg := config.Cfg.MySql["default"]

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.UserName,
		mysqlCfg.Password,
		mysqlCfg.DbHost,
		mysqlCfg.DbPort,
		mysqlCfg.DbName,
	)

	var ormLogger = logger.Default
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger, // 打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表明不加s
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置连接池中的最大闲置连接
	sqlDB.SetMaxOpenConns(100)                 // 设置数据库的最大连接数量
	sqlDB.SetConnMaxLifetime(30 * time.Second) // 设置连接的最大可复用时间

	_db = db

	// MakeMigrate(_db)

}

// MakeMigrate 迁移数据表，在没有数据表结构变更时候，建议注释不执行
// 只支持创建表、增加表中没有的字段和索引
// 为了保护数据，并不支持改变已有的字段类型或删除未被使用的字段
func MakeMigrate(db *gorm.DB) {
	err := db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.User{}, &model.Task{})

	if err != nil {
		panic("gorm 迁移失败")
	} else {
		log.Println("gorm 迁移成功")
	}

}
