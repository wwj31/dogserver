package common

import (
	"server/common/log"
	"server/proto/innermsg/inner"
	"server/proto/outermsg/outer"
)

func IsErr(v any, err error) (bool, outer.ERROR) {
	if err != nil {
		log.Errorw("assert failed ", "err", err)
		return true, outer.ERROR_FAILED
	}

	if errInfo, ok := v.(*inner.Error); ok {
		log.Errorw("assert failed with inner.Error", "inner.Error", errInfo.ErrorInfo)
		return true, outer.ERROR(errInfo.ErrorCode)
	}
	return false, 0
}
