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
		"shortId": 1744602,
		"gold":    10000,
	})

	req, err := http.NewRequest(http.MethodPost, "http://localhost:9999/alliance/setgold", bytes.NewReader(b))
	addSign(req)
	assert.Nil(t, err)

	c := &http.Client{}
	_, rspErr := c.Do(req)
	assert.Nil(t, rspErr)
}
