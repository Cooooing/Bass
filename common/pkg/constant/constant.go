package constant

import (
	v1 "common/api/common/v1"
	"math"
)

var page uint32 = 1
var size uint32 = 10

func GetPageDefault() *v1.PageRequest {
	return &v1.PageRequest{
		Page: page,
		Size: size,
	}
}
func GetPageMax() *v1.PageRequest {
	return &v1.PageRequest{
		Page: page,
		Size: math.MaxUint32,
	}
}
