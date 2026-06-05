package usecase

import (
	"content/internal/biz/repo"
	"context"
)

func accountName(ctx context.Context, userClient repo.UserClient, userID int64) (string, error) {
	if userID == 0 {
		return "", nil
	}
	accounts, err := userClient.MapAccounts(ctx, []int64{userID})
	if err != nil {
		return "", err
	}
	if accounts[userID] == nil {
		return "", nil
	}
	return accounts[userID].Name, nil
}
