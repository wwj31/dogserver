package gm_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// gm设置金币
func TestURLSetGold(t *testing.T) {
	rsp, err := http.Get("http://localhost:9999/gm/gold/?shortId=1587608&gold=100000")
	assert.Nil(t, err)
	fmt.Println(rsp)
}

// gm设置盟主
func TestGMSetMaster(t *testing.T) {
	rsp, err := http.Get("http://1.14.17.15:9999/gm/master/?shortId=1022696&rebate=99")
	assert.Nil(t, err)
	fmt.Println(rsp)
}

// gm设置盟主
func TestGMSetTime(t *testing.T) {
	rsp, err := http.Get("http://1.14.17.15:9999/gm/time/?date=2023-10-5 08:00:10")
	rsp, err = http.Get("http://1.14.17.15:9999/gm/time/?s=40")
	rsp, err = http.Get("http://1.14.17.15:9999/gm/time/?m=30")
	rsp, err = http.Get("http://1.14.17.15:9999/gm/time/?h=5")
	rsp, err = http.Get("http://1.14.17.15:9999/gm/time/?d=3")
	rsp, err = http.Get("http://1.14.17.15:9999/gm/cleartime")
	assert.Nil(t, err)
	fmt.Println(rsp)
}
