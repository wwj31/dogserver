FROM alpine:3.19

WORKDIR /server

RUN apk add tzdata # 时区相关的包,影响time.Now().Zone()是否能正常获得时区

COPY ../bin/dog .

ENTRYPOINT ["./dog"]