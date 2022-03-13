package localmsg

import "server/common"

type (
	Login struct {
		GSession common.GSession
		RId      uint64
		UId      uint64
	}
)
