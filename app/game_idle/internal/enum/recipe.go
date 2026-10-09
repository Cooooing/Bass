package enum

import (
	commonenum "common/pkg/enum"
	v1 "common/proto/gen/game_idle/v1/enum"
)

// RecipeType 表示配方类型；行动可以绑定多个不同类型的配方依次结算。
type RecipeType string

const (
	RecipeTypeNormal RecipeType = "normal" // 普通产出
	RecipeTypeRare   RecipeType = "rare"   // 稀有产出
)

var RecipeTypeMap = commonenum.NewMapping[RecipeType, v1.RecipeType](map[RecipeType]commonenum.Entry[RecipeType, v1.RecipeType]{
	RecipeTypeNormal: {Proto: v1.RecipeType_RECIPE_TYPE_NORMAL}, RecipeTypeRare: {Proto: v1.RecipeType_RECIPE_TYPE_RARE},
})

func (e RecipeType) String() string {
	return string(e)
}
