package common

var refactorTable = []func(){}

func RegistRefactorFun(f func()) {
	refactorTable = append(refactorTable, f)
}

// 配置表逻辑结构
func RefactorConfig() {
	for _, f := range refactorTable {
		f()
	}
}
