package iden

import "github.com/tiyee/palmon/consts"

func Score(prior consts.Prior, factor uint8) int64 {
	// 4+41
	factor %= 100
	maxScore := int64(1) << 62
	return maxScore - (int64(prior) << int64(41)) - int64(factor)*getMilliSeconds()
}
