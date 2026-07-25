package usecase

import (
	"context"
	"testing"

	"game_town/internal/biz/model"
	"game_town/internal/biz/repo"
)

type worldMemberRepoStub struct {
	repo.WorldMemberRepo
	member *model.WorldMember
}

func (r *worldMemberRepoStub) Get(
	context.Context,
	*repo.WorldMemberQuery,
) (*model.WorldMember, error) {
	return r.member, nil
}

type worldStateRepoStub struct {
	repo.WorldStateRepo
	state *model.WorldState
}

func (r *worldStateRepoStub) Get(
	context.Context,
	*repo.WorldStateQuery,
) (*model.WorldState, error) {
	return r.state, nil
}

func TestWorldMemberGet(
	t *testing.T,
) {
	usecase := NewWorldMemberUsecase(
		nil,
		nil,
		nil,
		&worldMemberRepoStub{
			member: &model.WorldMember{
				ID:       3,
				WorldID:  1,
				PlayerID: 2,
			},
		},
		&worldStateRepoStub{
			state: &model.WorldState{
				WorldID: 1,
				Version: 4,
			},
		},
		nil,
	)

	resp, err := usecase.Get(context.Background(), &GetWorldMemberReq{
		WorldID:  1,
		PlayerID: 2,
	})
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if resp.Member.ID != 3 || resp.State.Version != 4 {
		t.Fatalf("unexpected response: %#v", resp)
	}
}
