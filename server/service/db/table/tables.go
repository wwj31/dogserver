package table

var AllTable = make(map[string]Tabler)

type Tabler interface {
	TableName() string
	Count() int
	Key() uint64
}

func RegisterTable(table Tabler) {
	AllTable[table.TableName()] = table
}
