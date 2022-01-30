package gogotest

import (
	"fmt"
	"github.com/gogo/protobuf/proto"
	gogo "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"server/proto/innermsg/inner"
	"testing"
)

// gogo marshal to proto unmarshal
func TestGoGoMarshal_ProtoUnmarshal(t *testing.T) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	b, e := gogo.Marshal(&pt)
	assert.Empty(t, e)

	pt2 := inner.G2LRoleOffline{}
	assert.Empty(t, proto.Unmarshal(b, &pt2))
	fmt.Println(pt2.String())
}

// proto marshal to gogo unmarshal
func TestProtoUnmarshal_GoGoMarshal(t *testing.T) {
	pt := inner.G2LRoleOffline{}
	pt.RID = 3432423
	pt.UID = 12321321
	pt.GateSession = "this is test"
	b, e := proto.Marshal(&pt)
	assert.Empty(t, e)

	pt2 := inner.G2LRoleOffline{}
	assert.Empty(t, gogo.Unmarshal(b, &pt2))
	fmt.Println(pt2.String())
}
