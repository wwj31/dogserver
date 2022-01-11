package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"runtime/debug"
)

var (
	saveGoPath    = flag.String("saveGoPath", "./config", ".conf.go output path")
	saveJsonPath  = flag.String("saveJsonPath", "./config", ".json output path")
	readPath      = flag.String("readPath", "./", "The path of reading Excel")
	tplPath       = flag.String("tplPath", "./", "The path of .go.tpl")
	goPackageName = flag.String("goPackageName", "config", "go package of file")
)

func main() {
	flag.Parse()
	if *saveGoPath == "" || *readPath == "" || *goPackageName == "" || *saveJsonPath == "" || *tplPath == "" {
		fmt.Println("SaveGoPath, ReadPath or allType is nil")
		return
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover %v:%v", r, string(debug.Stack()))
		}
	}()
	fmt.Println("GoPath:")
	fmt.Println(filepath.Abs(*saveGoPath))

	fmt.Println("JsonPath:")
	fmt.Println(filepath.Abs(*saveJsonPath))

	fmt.Println("ExcelPath:")
	fmt.Println(filepath.Abs(*readPath))

	fmt.Println("TplPath:")
	fmt.Println(filepath.Abs(*tplPath))

	(&Generate{
		SaveGoPath:   *saveGoPath,
		SaveJsonPath: *saveJsonPath,
		ReadPath:     *readPath,
		PackageName:  *goPackageName,
		TplPath:      *tplPath,
	}).ReadExcel()
}
