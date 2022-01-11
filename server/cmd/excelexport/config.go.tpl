// Code generated by excelExoprt. DO NOT EDIT.
// source. {{.FileHeaderComment}}

package {{.Packagename}}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
)

// array and map
var _{{.TypeName}}Map = map[{{.KeyType}}]*{{.TypeName}}{}
var _{{.TypeName}}Array = []*{{.TypeName}}{}

type {{.TypeName}} struct {
    data *_{{.TypeName}}
}

// 类型结构
{{.StruType}}

func Has{{.TypeName}}(key {{.KeyType}}) bool {
	_, ok := _{{.TypeName}}Map[key]
	return ok
}

func Get{{.TypeName}}(key {{.KeyType}}) *{{.TypeName}} {
	return _{{.TypeName}}Map[key]
}

func Range{{.TypeName}}(fn func(i int,row *{{.TypeName}}) (stop bool))  {
	for i, row := range _{{.TypeName}}Array {
		if fn(i,row) {
        	break
       	}
	}
}

func Len{{.TypeName}}() int { return len(_{{.TypeName}}Array) }

func init() {
	loadFn["{{.TypeName}}"] = load{{.TypeName}}
}

func load{{.TypeName}}(dir string) error {
	data, err := ioutil.ReadFile(path.Join(dir, "{{.SheetName}}.json"))
	if err != nil {
		return fmt.Errorf("file=%v read err=%v", err.Error())
	}

	datas := []*_{{.TypeName}}{}
	err = json.Unmarshal(data, &datas)
	if err != nil {
		return fmt.Errorf("file=%v parse err=%v", err.Error())
	}

    result_array := make([]*{{.TypeName}},0,len(datas))
	result_map := make(map[{{.KeyType}}]*{{.TypeName}},len(datas))
	for _, row := range datas {
	    data := &{{.TypeName}}{data:row}
	    result_array = append(result_array, data)
		result_map[row.{{.Key}}] = data
	}
	_{{.TypeName}}Array = result_array
	_{{.TypeName}}Map = result_map
	fmt.Printf("%-50v len:%v\n", "{{.TypeName}} load finish! ", len(result_array))
	return nil
}
