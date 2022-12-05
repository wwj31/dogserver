package cache

import (
	"server/db/dbmysql/table"
)

type TableCache struct {
}

func (s *TableCache) Save(data table.Tabler) error {

	return nil
}

func (s *TableCache) Load(data table.Tabler) error {
	return nil
}

func (s *TableCache) LoadAll(tableName string, arr interface{}) error {
	return nil
}
