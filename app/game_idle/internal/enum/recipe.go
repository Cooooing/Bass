package enum

// RecipeType 表示配方类型；行动可以绑定多个不同类型的配方依次结算。
type RecipeType string

const (
	RecipeTypeNormal RecipeType = "normal" // 普通产出
	RecipeTypeRare   RecipeType = "rare"   // 稀有产出
)

func (e RecipeType) String() string {
	return string(e)
}

func RecipeTypeValues() []string {
	return []string{
		RecipeTypeNormal.String(),
		RecipeTypeRare.String(),
	}
}
