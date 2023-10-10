package req_test

import (
	"bytes"
	"encoding/json"
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
