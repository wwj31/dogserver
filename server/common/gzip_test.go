package common

import (
	"fmt"
	"testing"
)

func TestGZip(t *testing.T) {
	testData := []byte(
		`ABCDEFGHIJfjiewofjiewofjewiofjewiofjewiofjewiofjewiofjewiofjewofjew
		FJEWIOFJEWIOFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIOFJWEOIFJEWOIFJEWIOFJ
		FJEIOWFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIFOEWJFIOEWJFIOEWJFWEOIFWEIO
		VNREIGHREIGHREIWOGHREIOGHREIJWOGHREIWUOGHREUIWGHREUIWOGHRUIEWOGHRUEI
		VHUREOWTYRQPTFEUWOPFHDISVBNUERIWOGNVUIRONVURIEOGHRUEIWGHOREUITUQIOBN
		FJEWIOFJEWIOFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIOFJWEOIFJEWOIFJEWIOFJ
		FJEIOWFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIFOEWJFIOEWJFIOEWJFWEOIFWEIO
		VNREIGHREIGHREIWOGHREIOGHREIJWOGHREIWUOGHREUIWGHREUIWOGHRUIEWOGHRUEI
		"KLMN，这是测试数据`)
	fmt.Println("data size", len(testData))
	compress, _ := GZip(testData)
	fmt.Println(compress)
	fmt.Println("compress size", len(compress))
	data, _ := UnGZip(compress)
	fmt.Println(string(data))
}

var zipdata []byte

func init() {
	testData := []byte("abcdefe")
	zipdata, _ = GZip(testData)
}

func BenchmarkGZip(b *testing.B) {
	b.ReportAllocs()
	testData := []byte("abcdefe")
	for i := 0; i < b.N; i++ {
		_, _ = GZip(testData)
	}
}

func BenchmarkTestGZip(b *testing.B) {
	b.ReportAllocs()
	testData := []byte("abcdefe")
	for i := 0; i < b.N; i++ {
		_, _ = testGZip(testData)
	}
}

func BenchmarkUnGZip(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = UnGZip(zipdata)
	}
}
func BenchmarkTestUnGZip(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = testUnGZip(zipdata)
	}
}
