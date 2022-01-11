package common

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func GZip(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer, _ := gzip.NewWriterLevel(&buf, gzip.BestSpeed)
	_, err := writer.Write(data)
	if err != nil {
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func UnGZip(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err == nil {
		data, err = ioutil.ReadAll(gz)
		return data, nil
	}
	return nil, err
}
