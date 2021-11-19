package db

import (
	"database/sql"
	"fmt"
	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"server/service/db/table"
	"strconv"
	"time"
)

type DB struct {
	addr         string
	databaseAddr string
	dbIns        *gorm.DB
}

func New(addr, databaseAddr string) *DB {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(addr, databaseAddr)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	expect.Nil(err)

	// 检查库是否存在，不存在就创建
	expect.Nil(checkDatabase(fmt.Sprintf(addr, ""), databaseAddr))

	// 检查表是否存在，不存在就创建
	expect.Nil(checkTables(fmt.Sprintf(addr, databaseAddr), db))

	return &DB{
		addr:         addr,
		databaseAddr: databaseAddr,
		dbIns:        db,
	}
}

func (s *DB) Save(data table.Tabler) error {
	name := data.TableName()
	if data.Count() > 1 {
		name = name + strconv.Itoa(int(data.Key()%uint64(data.Count()))+1)
	}

	return s.dbIns.Table(name).Save(data).Error
}

func (s *DB) Load(data table.Tabler) error {
	name := data.TableName()
	if data.Count() > 1 {
		name = name + strconv.Itoa(int(data.Key()%uint64(data.Count()))+1)
	}
	return s.dbIns.Table(name).First(data).Error
}

func (s *DB) LoadAll(data table.Tabler, arr interface{}) error {
	name := data.TableName()
	if data.Count() > 1 {
		name = name + strconv.Itoa(int(data.Key()%uint64(data.Count()))+1)
	}
	return s.dbIns.Table(name).Find(arr).Error
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

func checkTables(addr string, db *gorm.DB) error {
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
		database := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4")
		if t.Count() <= 1 {
			creatErr = database.AutoMigrate(t)
			if creatErr != nil {
				return creatErr
			}
		} else {
			for i := 1; i <= t.Count(); i++ {
				creatErr = database.Table(t.TableName() + strconv.Itoa(i)).AutoMigrate(t)
				if creatErr != nil {
					return creatErr
				}
			}
		}
	}
	return nil
}
