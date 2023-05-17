// Code generated by excelExport. DO NOT EDIT.
// source. 测试2.xlsx

package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

// array and map
var _Tes2tMap = map[int64]*Tes2t{}
var _Tes2tArray = []*Tes2t{}

type Tes2t struct {
	data *_Tes2t
}

// 类型结构

type _Tes2t struct {
	Id     int64       // ID
	Test1  string      // 测试字段1
	Test2  int64       // 者是个
	Test4  int2int     // 飞机我
	Test5  int2str     // 分解为
	Test6  str2int     // 附件
	Test7  str2str     // 夫君
	Test8  array_int   // 额外i吧
	Test9  array_str   // 非飞机
	Test10 float32     // 飞机我附件二
	Test11 array_float // 附件
}

//ID
func (c *Tes2t) Id() int64 { return c.data.Id }

// 测试字段1
func (c *Tes2t) Test1() string { return c.data.Test1 }

// 者是个
func (c *Tes2t) Test2() int64 { return c.data.Test2 }

// 飞机我
func (c *Tes2t) LenTest4() int                                { return c.data.Test4.Len() }
func (c *Tes2t) Test4(key int64) (int64, bool)                { return c.data.Test4.Get(key) }
func (c *Tes2t) RangeTest4(fn func(int64, int64) (stop bool)) { c.data.Test4.Range(fn) }
func (c *Tes2t) CopyTest4() int2int                           { return c.data.Test4.Copy() }

// 分解为
func (c *Tes2t) LenTest5() int                                 { return c.data.Test5.Len() }
func (c *Tes2t) Test5(key int64) (string, bool)                { return c.data.Test5.Get(key) }
func (c *Tes2t) RangeTest5(fn func(int64, string) (stop bool)) { c.data.Test5.Range(fn) }
func (c *Tes2t) CopyTest5() int2str                            { return c.data.Test5.Copy() }

// 附件
func (c *Tes2t) LenTest6() int                                 { return c.data.Test6.Len() }
func (c *Tes2t) Test6(key string) (int64, bool)                { return c.data.Test6.Get(key) }
func (c *Tes2t) RangeTest6(fn func(string, int64) (stop bool)) { c.data.Test6.Range(fn) }
func (c *Tes2t) CopyTest6() str2int                            { return c.data.Test6.Copy() }

// 夫君
func (c *Tes2t) LenTest7() int                                  { return c.data.Test7.Len() }
func (c *Tes2t) Test7(key string) (string, bool)                { return c.data.Test7.Get(key) }
func (c *Tes2t) RangeTest7(fn func(string, string) (stop bool)) { c.data.Test7.Range(fn) }
func (c *Tes2t) CopyTest7() str2str                             { return c.data.Test7.Copy() }

// 额外i吧
func (c *Tes2t) LenTest8() int                              { return c.data.Test8.Len() }
func (c *Tes2t) Test8(key int) (int64, bool)                { return c.data.Test8.Get(key) }
func (c *Tes2t) RangeTest8(fn func(int, int64) (stop bool)) { c.data.Test8.Range(fn) }
func (c *Tes2t) CopyTest8() array_int                       { return c.data.Test8.Copy() }

// 非飞机
func (c *Tes2t) LenTest9() int                               { return c.data.Test9.Len() }
func (c *Tes2t) Test9(key int) (string, bool)                { return c.data.Test9.Get(key) }
func (c *Tes2t) RangeTest9(fn func(int, string) (stop bool)) { c.data.Test9.Range(fn) }
func (c *Tes2t) CopyTest9() array_str                        { return c.data.Test9.Copy() }

// 飞机我附件二
func (c *Tes2t) Test10() float32 { return c.data.Test10 }

// 附件
func (c *Tes2t) LenTest11() int                                { return c.data.Test11.Len() }
func (c *Tes2t) Test11(key int) (float32, bool)                { return c.data.Test11.Get(key) }
func (c *Tes2t) RangeTest11(fn func(int, float32) (stop bool)) { c.data.Test11.Range(fn) }
func (c *Tes2t) CopyTest11() array_float                       { return c.data.Test11.Copy() }

func HasTes2t(key int64) bool {
	_, ok := _Tes2tMap[key]
	return ok
}

func GetTes2t(key int64) *Tes2t {
	return _Tes2tMap[key]
}

func RangeTes2t(fn func(i int, row *Tes2t) (stop bool)) {
	for i, row := range _Tes2tArray {
		if fn(i, row) {
			break
		}
	}
}

func LenTes2t() int { return len(_Tes2tArray) }

func init() {
	loadFn["Tes2t"] = loadTes2t
}

func loadTes2t(dir string) error {
	data, err := os.ReadFile(path.Join(dir, "Tes2t.json"))
	if err != nil {
		return fmt.Errorf("file=%v read err=%v", err.Error())
	}

	datas := []*_Tes2t{}
	err = json.Unmarshal(data, &datas)
	if err != nil {
		return fmt.Errorf("file=%v parse err=%v", err.Error())
	}

	result_array := make([]*Tes2t, 0, len(datas))
	result_map := make(map[int64]*Tes2t, len(datas))
	for _, row := range datas {
		data := &Tes2t{data: row}
		result_array = append(result_array, data)
		result_map[row.Id] = data
	}
	_Tes2tArray = result_array
	_Tes2tMap = result_map
	fmt.Printf("%-50v len:%v\n", "Tes2t load finish! ", len(result_array))
	return nil
}