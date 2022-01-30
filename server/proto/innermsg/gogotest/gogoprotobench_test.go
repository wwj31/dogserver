package gogotest

import (
	"github.com/gogo/protobuf/proto"
	gogo "github.com/gogo/protobuf/proto"
	"server/proto/innermsg/inner"
	"testing"
)

func BenchmarkGoGoMarshal(b *testing.B) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gogo.Marshal(&pt)
	}
}
func BenchmarkProtoMarshal(b *testing.B) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		proto.Marshal(&pt)
	}
}

func BenchmarkGoGoUnmarshal(b *testing.B) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	bytes, _ := gogo.Marshal(&pt)
	b.ReportAllocs()
	b.ResetTimer()

	pt2 := inner.G2LRoleOffline{}
	for i := 0; i < b.N; i++ {
		gogo.Unmarshal(bytes, &pt2)
	}
}
func BenchmarkProtoUnmarshal(b *testing.B) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	bytes, _ := proto.Marshal(&pt)
	b.ReportAllocs()
	b.ResetTimer()

	pt2 := inner.G2LRoleOffline{}
	for i := 0; i < b.N; i++ {
		proto.Unmarshal(bytes, &pt2)
	}
}
