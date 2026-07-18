package base

type PageRequest struct {
	Page int64
	Size int64
}

type PageResp struct {
	Page  int64
	Size  int64
	Total int64
}
