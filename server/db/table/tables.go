package table

var AllTable = make(map[string]Tabler)

type Tabler interface {
	ModelName() string
	Count() int
	Key() uint64
}

func RegisterTable(table Tabler) {
	AllTable[table.ModelName()] = table
}
