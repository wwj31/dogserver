package localmsg

import "server/common"

type (
	Login struct {
		GSession common.GSession
		RId      string
		UId      string
		First    bool
	}
)
