package db

import (
	"database/sql"
	"fmt"
	"server/db/table"
	"strconv"
	"strings"
	"time"

	"github.com/wwj31/dogactor/expect"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DataCenter struct {
	addr   string
	dbName string
	dbIns  *gorm.DB
}

func New(addr, databaseName string) *DataCenter {
	// 检查库是否存在，不存在就创建
	expect.Nil(checkDatabase(fmt.Sprintf(addr, ""), databaseName))

	db, err := gorm.Open(mysql.Open(fmt.Sprintf(addr, databaseName)), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	expect.Nil(err)

	// 创建不存在的表
	expect.Nil(checkTables(db))
	return &DataCenter{
		addr:   addr,
		dbName: databaseName,
		dbIns:  db,
	}
}

func (s *DataCenter) Store(insert bool, tablers ...table.Tabler) error {
	if len(tablers) == 0 {
		return nil
	}
	for _, t := range tablers {
		name := t.TableName()
		if t.Count() > 1 {
			name = name + strconv.Itoa(num(t.Key(), t.Count()))
		}

		// Save和Updates的区别 https://learnku.com/docs/gorm/v2/update/9734
		if insert {
			s.dbIns.Table(name).Save(t)
		} else {
			s.dbIns.Table(name).Updates(t)
		}
	}
	return nil
}

func (s *DataCenter) Load(tablers ...table.Tabler) error {
	for _, t := range tablers {
		name := t.TableName()
		if t.Count() > 1 {
			name = name + strconv.Itoa(num(t.Key(), t.Count()))
		}
		s.dbIns.Table(name).First(t)
	}
	return nil
}

func (s *DataCenter) LoadAll(tableName string, arr interface{}) error {
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
	return err
}

func checkTables(db *gorm.DB) error {
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
