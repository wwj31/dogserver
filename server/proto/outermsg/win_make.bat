protoc -I=.\proto --plugin protoc-gen-go=../exec/protoc-gen-go.exe --go_out=./ .\proto\*.proto

pause