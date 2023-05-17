mkdir pbjs
protoc -I=./proto --plugin protoc-gen-go=../exec/protoc-gen-js --go_out=./pbjs ./proto/*.proto
zip -r pbjs.zip pbjs/
rm -rf pbjs