package req_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

//const addr = "localhost"

const addr = "1.14.17.15"

func TestSetMaster(t *testing.T) {
	b, _ := json.Marshal(map[string]interface{}{
		"shortId": 1885298,
	})
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("http://%v:9999/alliance/setmaster", addr),
		bytes.NewReader(b))
	addSign(req)
	assert.Nil(t, err)

	c := &http.Client{}
	result, rspErr := c.Do(req)
	assert.Nil(t, rspErr)
	v, _ := io.ReadAll(result.Body)
	fmt.Println("resp ", string(v))
}
