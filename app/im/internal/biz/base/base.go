package base

type PageRequest struct {
	Page int64
	Size int64
}

type PageResponse struct {
	Page  int64
	Size  int64
	Total int64
}
