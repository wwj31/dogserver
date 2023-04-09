// Code generated by excelExport. DO NOT EDIT.
// source. 全局表.xlsx

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// array and map
var _GlobalMap = map[int64]*Global{}
var _GlobalArray = []*Global{}

type Global struct {
	data *_Global
}

// 类型结构

type _Global struct {
	Id  int64  // ID
	Val string // 值
}

// ID
func (c *Global) Id() int64 { return c.data.Id }

// 值
func (c *Global) Val() string { return c.data.Val }

func HasGlobal(key int64) bool {
	_, ok := _GlobalMap[key]
	return ok
}

func GetGlobal(key int64) *Global {
	return _GlobalMap[key]
}

func RangeGlobal(fn func(i int, row *Global) (stop bool)) {
	for i, row := range _GlobalArray {
		if fn(i, row) {
			break
		}
	}
}

func LenGlobal() int { return len(_GlobalArray) }

func init() {
	loadFn["Global"] = loadGlobal
}

func loadGlobal(dir string) error {
	data, err := os.ReadFile(path.Join(dir, "global.json"))
	if err != nil {
		return fmt.Errorf("file=%v read err=%v", err.Error())
	}

	datas := []*_Global{}
	err = json.Unmarshal(data, &datas)
	if err != nil {
		return fmt.Errorf("file=%v parse err=%v", err.Error())
	}

	result_array := make([]*Global, 0, len(datas))
	result_map := make(map[int64]*Global, len(datas))
	for _, row := range datas {
		data := &Global{data: row}
		result_array = append(result_array, data)
		result_map[row.Id] = data
	}
	_GlobalArray = result_array
	_GlobalMap = result_map
	fmt.Printf("%-50v len:%v\n", "Global load finish! ", len(result_array))
	return nil
}
