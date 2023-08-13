package req_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetGold(t *testing.T) {
	b, _ := json.Marshal(map[string]interface{}{
		"shortId": 1905995,
		"gold":    1000,
	})

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/alliance/addgold", bytes.NewReader(b))
	addSign(req)
	assert.Nil(t, err)

	c := &http.Client{}
	_, rspErr := c.Do(req)
	assert.Nil(t, rspErr)
}

func TestURLSetGold(t *testing.T) {
	rsp, err := http.Get("http://localhost:9999/gm/gold/?shortId=1639901&gold=1000")
	assert.Nil(t, err)
	fmt.Println(rsp)
}
