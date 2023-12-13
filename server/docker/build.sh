cd ..

if [[ "$OS" == "Windows_NT" ]]; then
  localPath=/$(pwd)
else
  localPath=$(pwd)
fi

docker run --rm -t -v "$localPath:/server" -v "$GOPATH/pkg/mod:/root/go/pkg/mod" wwj32/alpine-go1.21 sh -c "cd server/cmd/main && go build -o ../../docker/bin/dog ."

cd docker
docker build -t "dog" .
docker tag dog wwj32/dog:latest
docker push wwj32/dog:latest

rm -rf bin