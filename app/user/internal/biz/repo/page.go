package repo

type PageReq struct {
	Page uint32
	Size uint32
}

type PageResp struct {
	Total uint32
	Page  uint32
	Size  uint32
}
