package common

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"sync"
	"unsafe"
)

var (
	bytesBufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}
	zipWriterPool   = sync.Pool{New: func() interface{} { return new(gzip.Writer) }}
	zipReaderPool   = sync.Pool{New: func() interface{} { return new(gzip.Reader) }}
)

func GZip(data []byte) ([]byte, error) {
	buf := bytesBufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bytesBufferPool.Put(buf)
	}()

	zip := zipWriterPool.Get().(*gzip.Writer)
	defer zipWriterPool.Put(zip)

	zip.Reset(buf)
	_, err := zip.Write(data)
	if err != nil {
		return nil, err
	}
	err = zip.Close()
	if err != nil {
		return nil, err
	}

	zipData := make([]byte, buf.Len())
	copy(zipData, buf.Bytes())
	return zipData, nil
}

// testGZip to compared with GZip for benchmark
func testGZip(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	w := gzip.NewWriter(buf)
	err := w.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnGZip(data []byte) ([]byte, error) {
	zip := zipReaderPool.Get().(*gzip.Reader)
	defer zipReaderPool.Put(zip)

	buf := bytesBufferPool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		bytesBufferPool.Put(buf)
	}()

	// use to assign buf of bytes.Buffer
	ptr := unsafe.Pointer(buf)
	t := (*struct{ buf []byte })(ptr)
	t.buf = data

	err := zip.Reset(buf)
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zip)
	return data, err
}

// testUnGZip to compared with UnGZip for benchmark
func testUnGZip(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(gz)
}
func testUnGZip2(data []byte) ([]byte, error) {
	zip := zipReaderPool.Get().(*gzip.Reader)
	defer zipReaderPool.Put(zip)
	err := zip.Reset(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	data, err = ioutil.ReadAll(zip)
	return data, err
}
