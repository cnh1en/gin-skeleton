package initialize

import (
	"fmt"
	"time"

	"github.com/cnh1en/lets_go/global"
	"github.com/cnh1en/lets_go/internal/models"
	"github.com/cnh1en/lets_go/pkg/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMysql() {
	c := global.Config.Mysql

	f := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(f, c.Username, c.Password, c.Host, c.Port, c.Dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		utils.CheckErrorPanic(err, "init mysql")
	}
	global.DB = db
	global.Logger.Info("init mysql success")

	SetPool()
	migrateTables()
}

func SetPool() {
	m := global.Config.Mysql
	sqlDB, err := global.DB.DB()
	if err != nil {
		utils.CheckErrorPanic(err, "set pool")
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDB.SetMaxOpenConns(m.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))

}

func migrateTables() {
	err := global.DB.AutoMigrate(
		&entities.User{},
		&entities.Role{},
	)

	utils.CheckErrorPanic(err, "migrate tables")
}
