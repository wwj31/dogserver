cd ../../

docker build -t wwj32/alpine-go1.21:latest -f docker/build-alpine-go/Dockerfile .
docker push wwj32/alpine-go1.21:latest