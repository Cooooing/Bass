package repo

type PageReq struct {
	Page uint32
	Size uint32
}

type PageResponse struct {
	Page  uint32
	Size  uint32
	Total uint32
}
