package constant

import (
	v1 "common/api/common/v1"
	"math"
)

var page uint32 = 1
var size uint32 = 10

func PageValid(p *v1.PageRequest) *v1.PageRequest {
	if p == nil {
		return GetPageDefault()
	}
	if p.Page <= 0 {
		p.Page = page
	}
	if p.Size <= 0 {
		p.Size = size
	}
	return p
}
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
