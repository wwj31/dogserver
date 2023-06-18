package common

import (
	"encoding/json"
	"server/common/log"
)

func JsonMarshal(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Errorw("json marshal failed", "v", v)
		return ""
	}
	return string(b)
}

func JsonUnmarshal(str string, v any) {
	err := json.Unmarshal([]byte(str), v)
	if err != nil {
		log.Errorw("json unmarshal failed", "err", err)
	}
}
