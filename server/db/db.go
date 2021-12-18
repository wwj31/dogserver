package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"server/db/table"
	"strconv"
	"strings"
	"time"

	"github.com/wwj31/dogactor/expect"
	"github.com/wwj31/dogactor/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	addr         string
	databaseAddr string
	dbIns        *gorm.DB
}

func New(addr, databaseAddr string) *DB {
	// 检查库是否存在，不存在就创建
	expect.Nil(checkDatabase(fmt.Sprintf(addr, ""), databaseAddr))

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(addr, databaseAddr)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	expect.Nil(err)

	// 检查表是否存在，不存在就创建
	expect.Nil(checkTables(fmt.Sprintf(addr, databaseAddr), db))

	return &DB{
		addr:         addr,
		databaseAddr: databaseAddr,
		dbIns:        db,
	}
}

func (s *DB) Save(datas ...table.Tabler) error {
	if datas == nil {
		return nil
	}
	transaction := s.dbIns.Begin()
	for _, t := range datas {
		if reflect.ValueOf(t).IsNil() {
			continue
		}

		name := t.TableName()
		if t.Count() > 1 {
			name = name + strconv.Itoa(num(t.Key(), t.Count()))
		}
		transaction.Table(name).Save(t)
	}

	err := transaction.Commit().Error
	if err != nil {
		transaction.Rollback()
	}

	return err
}

func (s *DB) Load(datas ...table.Tabler) error {
	transaction := s.dbIns.Begin()
	for _, t := range datas {
		name := t.TableName()
		if t.Count() > 1 {
			name = name + strconv.Itoa(num(t.Key(), t.Count()))
		}
		transaction.Table(name).First(t)
	}

	return transaction.Commit().Error
}

func (s *DB) LoadAll(tableName string, arr interface{}) error {
	tableName = strings.ToLower(tableName)
	return s.dbIns.Table(tableName).Find(arr).Error
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

func num(key uint64, count int) int {
	v := int((key>>32)%uint64(count)) + 1
	return v
}
