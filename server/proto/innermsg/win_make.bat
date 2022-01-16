..\exec\protoc --plugin protoc-gen-go=..\exec\protoc-gen-go.exe -I=./ --go_out=./ .\proto\*.proto
go generate ./inner/type.go
pause