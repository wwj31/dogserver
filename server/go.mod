module server

go 1.15

replace (
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/golang/protobuf v1.5.0
	github.com/google/uuid v1.2.0 // indirect
	github.com/spf13/cast v1.4.1
	github.com/wwj31/dogactor v1.0.4
	github.com/yuin/gopher-lua v0.0.0-20200816102855-ee81675732da
	google.golang.org/protobuf v1.26.0
)
