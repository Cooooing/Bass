package repo

import "game_town/internal/biz/base"

type pageHelper struct{}

func (pageHelper) page(req base.PageRequest) base.PageRequest {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.Size < 1 {
		req.Size = 20
	}
	if req.Size > 100 {
		req.Size = 100
	}
	return req
}

func (pageHelper) pageOffset(req base.PageRequest) int {
	return int((req.Page - 1) * req.Size)
}

func (pageHelper) pageLimit(req base.PageRequest) int {
	return int(req.Size)
}

func (pageHelper) basePage(total int, req base.PageRequest) base.PageResp {
	return base.PageResp{
		Page:  req.Page,
		Size:  req.Size,
		Total: int64(total),
	}
}
