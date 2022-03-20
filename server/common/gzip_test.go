package common

import (
	"fmt"
	"testing"
)

var testData = []byte(
	`ABCDEFGHIJfjiewofjiewofjewiofjewiofjewiofjewiofjewiofjewiofjewofjew
		FJEWIOFJEWIOFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIOFJWEOIFJEWOIFJEWIOFJ
		FJEIOWFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIFOEWJFIOEWJFIOEWJFWEOIFWEIO
		VNREIGHREIGHREIWOGHREIOGHREIJWOGHREIWUOGHREUIWGHREUIWOGHRUIEWOGHRUEI
		VHUREOWTYRQPTFEUWOPFHDISVBNUERIWOGNVUIRONVURIEOGHRUEIWGHOREUITUQIOBN
		FJEWIOFJEWIOFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIOFJWEOIFJEWOIFJEWIOFJ
		FJEIOWFJEWIOFJEWIOFJEWIOFJEWOIFJEWIOFJEWIFOEWJFIOEWJFIOEWJFWEOIFWEIO
		VNREIGHREIGHREIWOGHREIOGHREIJWOGHREIWUOGHREUIWGHREUIWOGHRUIEWOGHRUEI
		"KLMN，这是测试数据`)

func TestGZip(t *testing.T) {
	fmt.Println("data size", len(testData))
	compress, _ := GZip(testData)
	fmt.Println(compress)
	fmt.Println("compress size", len(compress))
	data, _ := UnGZip(compress)
	fmt.Println(string(data))
}

var zipdata []byte

func init() {
	zipdata, _ = GZip(testData)
}

func BenchmarkGZip(b *testing.B) {
	b.ReportAllocs()
	data := []byte("abcdefe")
	for i := 0; i < b.N; i++ {
		_, _ = GZip(data)
	}
}

func BenchmarkTestGZip(b *testing.B) {
	b.ReportAllocs()
	data := []byte("abcdefe")
	for i := 0; i < b.N; i++ {
		_, _ = testGZip(data)
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

func BenchmarkTestUnGZip2(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = testUnGZip2(zipdata)
	}
}
