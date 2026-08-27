package model

// Region 是前端展示区域配置，不参与行动调度。
type Region struct {
	ID          string
	Name        string
	Description string
	Enabled     bool
	Sort        int32
}
