package service

import (
	"common/pkg/apperror"
	cerrors "common/proto/gen/common/errors"
	v1 "common/proto/gen/game_idle/v1"
	"context"
	"game_idle/internal/biz/model"
	"game_idle/internal/biz/usecase"

	"github.com/go-kratos/kratos/v3/transport/grpc"
	"github.com/go-kratos/kratos/v3/transport/http"
)

type ActionService struct {
	v1.UnimplementedActionServiceServer
	actionUsecase *usecase.ActionUsecase
}

func NewActionService(actionUsecase *usecase.ActionUsecase) *ActionService {
	return &ActionService{
		actionUsecase: actionUsecase,
	}
}

func (s *ActionService) RegisterGrpc(server *grpc.Server) {
	v1.RegisterActionServiceServer(server, s)
}

func (s *ActionService) RegisterHttp(*http.Server) {
}

func (s *ActionService) List(ctx context.Context, req *v1.ListActions_Request) (*v1.ListActions_Resp, error) {
	rows, err := s.actionUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	out := make([]*v1.Action, 0, len(rows))
	for _, row := range rows {
		out = append(out, s.buildAction(row))
	}
	return &v1.ListActions_Resp{Rows: out}, nil
}

func (s *ActionService) GetDetail(ctx context.Context, req *v1.GetActionDetail_Request) (*v1.GetActionDetail_Resp, error) {
	if req.GetActionId() == "" {
		return nil, apperror.New(cerrors.BusinessErrorCode_BUSINESS_ERROR_CODE_COMMON_INVALID_ARGUMENT)
	}
	row, err := s.actionUsecase.GetDetail(ctx, req.GetActionId())
	if err != nil {
		return nil, err
	}
	out := &v1.ActionDetail{
		Action:  s.buildAction(row.Action),
		Recipes: make([]*v1.ActionRecipeDetail, 0, len(row.Recipes)),
	}
	for _, recipe := range row.Recipes {
		recipeRow := &v1.ActionRecipeDetail{
			RecipeId:        recipe.ID,
			Name:            recipe.Name,
			Description:     recipe.Description,
			RecipeType:      recipe.Type.String(),
			GenerationTimes: recipe.GenerationTimes,
			Inputs:          make([]*v1.RecipeInputDetail, 0, len(recipe.Inputs)),
			Outputs:         make([]*v1.RecipeOutputDetail, 0, len(recipe.Outputs)),
		}
		for _, input := range recipe.Inputs {
			item := row.Items[input.ItemID]
			recipeRow.Inputs = append(recipeRow.Inputs, &v1.RecipeInputDetail{
				ItemId:   input.ItemID,
				ItemName: item.Name,
				ItemType: item.Type.String(),
				Quantity: input.Quantity,
			})
		}
		for _, output := range recipe.Outputs {
			item := row.Items[output.ItemID]
			probability := 0.0
			if recipe.TotalWeight > 0 {
				probability = float64(output.Weight) / float64(recipe.TotalWeight)
			}
			recipeRow.Outputs = append(recipeRow.Outputs, &v1.RecipeOutputDetail{
				ItemId:      output.ItemID,
				ItemName:    item.Name,
				ItemType:    item.Type.String(),
				MinQuantity: output.MinQuantity,
				MaxQuantity: output.MaxQuantity,
				Weight:      output.Weight,
				Probability: probability,
			})
		}
		out.Recipes = append(out.Recipes, recipeRow)
	}
	return &v1.GetActionDetail_Resp{Row: out}, nil
}

func (s *ActionService) buildAction(row *model.Action) *v1.Action {
	return &v1.Action{
		ActionId:             row.ID,
		Name:                 row.Name,
		Description:          row.Description,
		RegionId:             row.RegionID,
		ActionKind:           row.ActionKind.String(),
		AbilityId:            row.AbilityID,
		RequiredAbilityLevel: row.RequiredAbilityLevel,
		DurationSeconds:      int64(row.Duration.Seconds()),
		ExpReward:            row.ExpReward,
		Enabled:              row.Enabled,
		Sort:                 row.Sort,
	}
}
