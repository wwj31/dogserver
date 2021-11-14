package db

import (
	"database/sql"
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/iniconfig"
	"github.com/wwj31/dogactor/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/service/db/table"
	"time"
)

type DB struct {
	addr   string
	Config iniconfig.Config
}

func New(conf iniconfig.Config) *DB {
	addr := conf.String("mysql") // root:123456@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local
	database := conf.String("database")

	// 检查库是否存在，不存在就创建
	expect.Nil(checkDatabase(fmt.Sprintf(addr, ""), database))

	// 检查表是否存在，不存在就创建
	expect.Nil(checkTables(fmt.Sprintf(addr, database)))

	return &DB{
		addr:   addr,
		Config: conf,
	}
}

func (db *DB) Save(key string, data interface{}) error {
	return nil
}
func (db *DB) Load(key string) interface{} {
	return nil
}

func checkDatabase(addr string, databaseName string) error {
	msql, err := sql.Open("mysql", addr)
	if err != nil {
		return err
	}
	defer msql.Close()

	sql_format := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci", databaseName)
	_, err = msql.Exec(sql_format)
	if err == nil {
		log.KVs(log.Fields{"table": databaseName}).Debug("database create success")
	} else {
		return err
	}
	return nil
}

func checkTables(addr string) error {
	db, err := gorm.Open(mysql.Open(addr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		return err
	}

	db.Logger.LogMode(logger.Info)

	sqlDB, errdb := db.DB()
	if errdb != nil {
		return errdb
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 10)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	var creatErr error
	for _, t := range table.AllTable {
		if t.Count() <= 1 {
			creatErr = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
				AutoMigrate(t)
			if creatErr != nil {
				return creatErr
			}
		} else {
			for i := 1; i < t.Count(); i++ {
				creatErr = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").
					AutoMigrate(t)
				if creatErr != nil {
					return creatErr
				}
			}
		}
	}
	return nil
}
