package mysql

import (
	"clms/models"
	"clms/settings"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var dbs *gorm.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	// 也可以使用MustConnect连接不成功就panic
	dbs, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//解决查表是，会自动添加复数的问题
			SingularTable: true,
		},
	})
	if err != nil {
		zap.L().Error("connect DB failed", zap.Error(err))
		return
	}
	sqlDB, err := dbs.DB()
	//设置空闲链接池中连接的最大数量
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	err = dbs.AutoMigrate(&models.User{},
		&models.Product{},
		&models.Carousel{},
		&models.Category{},
		&models.Favorite{},
		&models.ProductImg{},
		&models.Order{},
		&models.Cart{},
		&models.Admin{},
		&models.Address{},
		&models.Notice{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	return
}
