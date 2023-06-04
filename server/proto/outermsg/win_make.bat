go run github.com/wwj31/msgidgen@latest -pack=outer -path=./proto -tag=./outer -upper=false -prefix=Id
..\exec\protoc.exe -I=.\proto --plugin protoc-gen-go=../exec/protoc-gen-go.exe --go_out=./ .\proto\*.proto
go generate ./outer/type.go
pause