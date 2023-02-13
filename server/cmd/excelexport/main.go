package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

var (
	saveGoPath    = flag.String("saveGoPath", "./config", ".conf.go output path")
	saveJsonPath  = flag.String("saveJsonPath", "./config", ".json output path")
	inputPath     = flag.String("inputPath", "./", "The path of reading Excel")
	tplPath       = flag.String("tplPath", "./", "The path of .go.tpl")
	goPackageName = flag.String("goPackageName", "config", "go package of file")
)

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover %v:%v", r, string(debug.Stack()))
		}
	}()

	(&Generate{
		SaveGoPath:   *saveGoPath,
		SaveJsonPath: *saveJsonPath,
		ReadPath:     *inputPath,
		PackageName:  *goPackageName,
		TplPath:      *tplPath,
	}).ReadExcel()
}
