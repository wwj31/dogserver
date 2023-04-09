package main

import (
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"html/template"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
)

var (
	lineNumber   = 3        // 每个工作表需要读取的行数
	headerFromat = "%v\n\r" // 文件头

	structContext = `
	type _%v struct {
		%v
	}

    %v
	`
)

type Generate struct {
	SaveGoPath   string // 生成go文件的保存路径
	SaveJsonPath string // 生成json文件的保存路径
	TplPath      string // 模板文件路径
	ReadPath     string // excel表路径
	PackageName  string // 包名
}

// ReadExcel 读取excel
func (s *Generate) ReadExcel() {
	files, err := os.ReadDir(s.ReadPath)
	if err != nil {
		panic(fmt.Errorf("excel文件路径读取失败 此路径无效:%v error:%v", s.ReadPath, err))
	}

	pChNum := make(chan struct{}, 10)
	for i, file := range files {
		//if hasChinese(file.Name()) {
		//	continue
		//}
		if strings.Contains(file.Name(), "~$") {
			continue
		}
		if !strings.Contains(file.Name(), ".xlsx") {
			continue
		}

		pChNum <- struct{}{}
		go func(idx int) {
			defer func() {
				<-pChNum
			}()

			dir := path.Join(s.ReadPath, files[idx].Name())
			f, err := xlsx.OpenFile(dir)
			if err != nil {
				panic(fmt.Errorf("excel文件读取失败 无效文件:%v error:%v", dir, err))
			}

			// 遍历工作表
			for _, sheet := range f.Sheets {
				fileName := files[idx].Name()
				if err := s.BuildTypeStruct(sheet, fileName); err != nil {
					panic(err)
				}

				if err := s.BuildJsonStruct(sheet, fileName); err != nil {
					panic(err)
				}

				fmt.Printf("%v  %-15v ok.\n", files[idx].Name(), sheet.Name)
			}

		}(i)
	}

	// 等待所有表处理完成
	for len(pChNum) > 0 {
		runtime.Gosched()
	}

	// 导出依赖的字段解析函数
	tpath := path.Join(s.TplPath, "convert.go.tpl")
	tmpl, err := template.ParseFiles(tpath)
	if err != nil {
		panic(fmt.Errorf("模板文件读取失败，无效路径:%v err:%v", tpath, err))
	}

	newFile := path.Join(s.SaveGoPath, "convert.funtype.go")
	file, err := os.OpenFile(newFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		panic(fmt.Errorf("配置文件创建失败 无效的文件或路径:%v err:%v", newFile, err))
	}

	err = tmpl.Execute(file, struct {
		PackageName string
	}{s.PackageName})

	if err != nil {
		panic(fmt.Errorf("模板执行输出失败 err:%v", err))
	}
	exec.Command("go", "fmt", s.SaveGoPath).Run()
}

// BuildTypeStruct 构建类型结构
func (s *Generate) BuildTypeStruct(sheet *xlsx.Sheet, fileName string) error {
	sheetMetas := make([][]string, 0)
	// 判断表格中内容的行数是否小于需要读取的行数
	if sheet.MaxRow < lineNumber {
		return fmt.Errorf("ReadExcel sheet.MaxRow:%d < lineNumber:%d file:%v", sheet.MaxRow, lineNumber, fileName)
	}
	// 遍历列
	for i := 0; i < sheet.MaxCol; i++ {
		// 没有字段，不解析
		if strings.TrimSpace(sheet.Cell(FIELDNAME, i).Value) == "" {
			continue
		}
		var meta []string
		// 遍历行
		for j := 0; j < lineNumber; j++ {
			meta = append(meta, strings.TrimSpace(sheet.Cell(j, i).Value))
		}
		sheetMetas = append(sheetMetas, meta)
	}
	structType, err := parsing(sheetMetas, sheet.Name)
	if err != nil {
		return fmt.Errorf("fileName:\"%v\" is err:%v", fileName, err)
	}

	if structType == "" {
		return fmt.Errorf("ReadExcel s.data is nil")
	}
	structType += "\n"

	fieldName := firstRuneToUpper(strings.TrimSpace(sheet.Cell(FIELDNAME, 0).Value))

	var (
		keyType string
		ok      bool
	)

	fieldType := strings.TrimSpace(sheet.Cell(FIELDTYPE, 0).Value)
	if fieldType != STR && fieldType != INT {
		return fmt.Errorf("主键类型错误:%v %v %v", fieldType, fileName, sheet.Name)
	}

	if keyType, ok = TypeIndex[fieldType]; !ok {
		return fmt.Errorf("主键类型找不到:%v %v %v", fieldType, fileName, sheet.Name)
	}

	err = s.writeGolangFile(structType, sheet.Name, keyType, fieldName, fileName)
	if err != nil {
		return err
	}
	return nil
}

// BuildJsonStruct 构建json结构
func (s *Generate) BuildJsonStruct(sheet *xlsx.Sheet, fileName string) error {
	var array []string
	// 判断表格中内容的行数是否小于需要读取的行数
	dataLen := sheet.MaxRow - lineNumber
	if dataLen < 0 {
		return fmt.Errorf("ReadExcel dataLen < 0 dataLen:%v MaxRow:%v lineNumber:%v ", dataLen, sheet.MaxRow, lineNumber)
	}

	// 遍历列
	var err error
	checkUnique := make(map[string]struct{}, sheet.MaxRow)
	for i := lineNumber; i <= sheet.MaxRow; i++ {
		// 遍历行
		cell := sheet.Cell(i, 0)
		if cell == nil {
			return fmt.Errorf("fileName:%v cell == nil ", fileName)
		}

		// 单独处理key
		key := strings.TrimSpace(cell.Value)
		if key == "" {
			break
		}
		if _, ex := checkUnique[key]; ex {
			return fmt.Errorf("表:[%v] 页:[%v] 主键重复 key:[%v]", fileName, sheet.Name, key)
		}
		m := map[string]interface{}{}
		for j := 0; j < sheet.MaxCol; j++ {
			if strings.TrimSpace(sheet.Cell(1, j).Value) == "" {
				continue
			}
			fieldName := strings.TrimSpace(sheet.Cell(FIELDNAME, j).Value)
			fieldType := strings.TrimSpace(sheet.Cell(FIELDTYPE, j).Value)

			if m[fieldName], err = TypeConvert[fieldType](strings.TrimSpace(sheet.Cell(i, j).Value)); err != nil {
				return fmt.Errorf("TypeConvert error=%v i=%v  j=%v name=%v value=%v file=%v",
					err, i, j, fieldName, sheet.Cell(i, j).Value, fileName)
			}
		}
		if data, err := json.Marshal(m); err == nil {
			array = append(array, string(data))
		} else {
			return fmt.Errorf("json.Marshal(array) error:%v ", err)
		}
		checkUnique[key] = struct{}{}
	}

	if err := s.writeJsonFile("[\n    "+strings.Join(array, ",\n    ")+"\n]", sheet.Name); err != nil {
		return err
	}
	return nil
}

func (s *Generate) BuildConfigConst(sheet *xlsx.Sheet) error {
	// 遍历列 找到STR_SERVER_CONST
	constData := ""
	for col := 0; col < sheet.MaxCol; col++ {
		if sheet.Cell(FIELDNAME, col).Value == "STR_SERVER_CONST" {
			// 遍历行
			for row := lineNumber; row < sheet.MaxRow; row++ {
				cell := sheet.Cell(row, col)
				if cell == nil {
					return fmt.Errorf("sheetName=%s, cell == nil ", sheet.Name)
				}
				key := sheet.Cell(row, 0).Value
				if strings.TrimSpace(key) == "" {
					break
				}
				if cell.Value != "" {
					constData += cell.Value + " = " + key + "\n"
				}
			}
			constData = "\n const( \n" + constData + ")"
			break
		}
	}
	if constData == "" {
		return nil
	}
	f, err := os.OpenFile(path.Join(s.SaveGoPath, sheet.Name+".conf.go"), os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("BuildConfigConst Open is err:%v", err)
	}
	defer f.Close()
	_, err = f.WriteString(constData)
	if err != nil {
		return fmt.Errorf("BuildConfigConst WriteString is err:%v", err)
	}
	fmt.Printf("%-13v %v \n", "BuildConfigConst", sheet.Name)
	return nil
}
