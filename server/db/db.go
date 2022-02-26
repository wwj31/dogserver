package db

import (
	"database/sql"
	"fmt"
	"server/common"
	"server/db/table"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cast"
	"github.com/wwj31/dogactor/actor"
	"github.com/wwj31/dogactor/expect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const process_count = 4
const process = "process_"

type DataCenter struct {
	addr      string
	dbName    string
	dbIns     *gorm.DB
	sys       *actor.System
	processId []common.ActorId
}

func New(addr, databaseName string, sys *actor.System) *DataCenter {
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

	// 创建processor
	var ids []common.ActorId
	for i := 1; i <= process_count; i++ {
		processActor := actor.New(process+cast.ToString(i), &processor{session: db.Session(&gorm.Session{})})
		expect.Nil(sys.Add(processActor))
		ids = append(ids)
	}

	return &DataCenter{
		addr:      addr,
		dbName:    databaseName,
		dbIns:     db,
		processId: ids,
		sys:       sys,
	}
}

func (s *DataCenter) Store(insert bool, tablers ...table.Tabler) error {
	if len(tablers) == 0 {
		return nil
	}
	for _, t := range tablers {
		var state op
		if insert {
			state = _INSERT
		} else {
			state = _UPDATE
		}

		opera := operator{
			state: state,
			tab:   t,
		}

		err := s.sys.Send("", processId(t), "", opera)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *DataCenter) Load(t table.Tabler) error {
	f := make(chan struct{}, 1)
	opera := operator{
		state:  _LOAD,
		tab:    t,
		finish: f,
	}

	err := s.sys.Send("", processId(t), "", opera)
	if err != nil {
		return err
	}

	timer := time.NewTimer(30 * time.Second)
	select {
	case <-f:
		return nil
	case <-timer.C:
		return fmt.Errorf("DataCenter load timeout", "table", t.ModelName(), "key", t.Key())
	}
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
		for i := 1; i <= t.Count(); i++ {
			creatErr = database.Table(tableName(t.ModelName(), i)).AutoMigrate(t)
			if creatErr != nil {
				return creatErr
			}
		}
	}
	return nil
}

func processId(t table.Tabler) string {
	hash := num(t.Key(), process_count)
	return process + cast.ToString(hash)
}

func tableName(name string, count int) string {
	if count == 0 {
		return name
	}
	return name + strconv.Itoa(count)
}

func num(key uint64, count int) int {
	v := int((key>>32)%uint64(count)) + 1
	return v
}
