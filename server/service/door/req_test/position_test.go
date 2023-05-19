package req_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPosition(t *testing.T) {
	b, _ := json.Marshal(map[string]interface{}{
		"shortId":  12345,
		"position": "大哥",
	})
	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/alliance/position", bytes.NewReader(b))
	addSign(req)
	assert.Nil(t, err)

	c := &http.Client{}
	_, rspErr := c.Do(req)
	assert.Nil(t, rspErr)
}
