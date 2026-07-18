package handler

import (
	"context"
	"notify/internal/biz/model"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"
)

type userClientHandler struct {
	userClient repo.UserClient
}

func (h *userClientHandler) loadBasic(ctx context.Context, userID int64) (*model.UserAccount, error) {
	if h.userClient == nil || userID == 0 {
		return nil, nil
	}
	usersResp, err := h.userClient.MapAccounts(ctx, []int64{userID})
	if err != nil {
		return nil, err
	}
	return usersResp[userID], nil
}

func (h *userClientHandler) loadAccounts(ctx context.Context, userIDs ...int64) (map[int64]*model.UserAccount, error) {
	if h.userClient == nil {
		return map[int64]*model.UserAccount{}, nil
	}
	seen := map[int64]struct{}{}
	ids := make([]int64, 0, len(userIDs))
	for _, userID := range userIDs {
		if userID == 0 {
			continue
		}
		if _, ok := seen[userID]; ok {
			continue
		}
		seen[userID] = struct{}{}
		ids = append(ids, userID)
	}
	if len(ids) == 0 {
		return map[int64]*model.UserAccount{}, nil
	}
	resp, err := h.userClient.MapAccounts(ctx, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *userClientHandler) templateUser(userID int64, user *model.UserAccount) model.TemplateUser {
	data := model.TemplateUser{ID: userID}
	if user == nil {
		return data
	}
	data.ID = user.ID
	if data.ID == 0 {
		data.ID = userID
	}
	data.Name = user.Name
	data.Nickname = user.Nickname
	return data
}

type contentClientHandler struct {
	contentClient repo.ContentClient
}

func (h *contentClientHandler) articleTemplateData(article *model.ContentArticle) model.TemplateArticle {
	if article == nil {
		return model.TemplateArticle{}
	}
	return model.TemplateArticle{
		ID:    article.ID,
		Title: article.Title,
		Author: model.TemplateUser{
			ID:       article.AuthorID,
			Name:     article.AuthorName,
			Nickname: article.AuthorNickname,
		},
	}
}

func (h *contentClientHandler) commentTemplateData(comment *model.ContentComment) model.TemplateComment {
	if comment == nil {
		return model.TemplateComment{}
	}
	return model.TemplateComment{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		Content:   comment.Content,
		User: model.TemplateUser{
			ID:       comment.UserID,
			Name:     comment.UserName,
			Nickname: comment.UserNickname,
		},
		ReplyUser: model.TemplateUser{
			ID:   comment.ReplyUserID,
			Name: comment.ReplyUserName,
		},
		Article: h.articleTemplateData(comment.Article),
	}
}

type articleActorHandler struct {
	userClientHandler
	contentClientHandler
}

func (h *articleActorHandler) build(ctx context.Context, eventID string, articleID int64, senderID int64) (*usecase.NotificationContext, error) {
	var article *model.ContentArticle
	if h.contentClient != nil && articleID != 0 {
		articleResp, err := h.contentClient.GetArticle(ctx, articleID)
		if err != nil {
			return nil, err
		}
		article = articleResp
	}
	users, err := h.loadAccounts(ctx, senderID)
	if err != nil {
		return nil, err
	}
	templateData := model.ArticleActorTemplateData{
		Article: h.articleTemplateData(article),
		Actor:   h.templateUser(senderID, users[senderID]),
	}

	recipients := make([]*usecase.NotificationRecipient, 0, 1)
	if article != nil && article.AuthorID != 0 && article.AuthorID != senderID {
		recipients = append(recipients, &usecase.NotificationRecipient{UserID: article.AuthorID})
	}
	return &usecase.NotificationContext{
		EventID:      eventID,
		TemplateData: templateData,
		Recipients:   recipients,
	}, nil
}
