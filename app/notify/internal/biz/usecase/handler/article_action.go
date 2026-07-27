package handler

import (
	"context"
	"notify/internal/biz/model"
	templatedata "notify/internal/biz/model/template_data"
	"notify/internal/biz/repo"
	"notify/internal/biz/usecase"

	"github.com/samber/lo"
)

type userClientHandler struct {
	userClient repo.UserAccountRepo
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
	ids := lo.Uniq(lo.Filter(userIDs, func(userID int64, _ int) bool {
		return userID != 0
	}))
	if len(ids) == 0 {
		return map[int64]*model.UserAccount{}, nil
	}
	resp, err := h.userClient.MapAccounts(ctx, ids)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *userClientHandler) templateUser(userID int64, user *model.UserAccount) templatedata.User {
	data := templatedata.User{
		ID: userID,
	}
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

func (h *contentClientHandler) articleTemplateData(article *model.ContentArticle) templatedata.Article {
	if article == nil {
		return templatedata.Article{}
	}
	return templatedata.Article{
		ID:    article.ID,
		Title: article.Title,
		Author: templatedata.User{
			ID:       article.AuthorID,
			Name:     article.AuthorName,
			Nickname: article.AuthorNickname,
		},
	}
}

func (h *contentClientHandler) commentTemplateData(comment *model.ContentComment) templatedata.Comment {
	if comment == nil {
		return templatedata.Comment{}
	}
	return templatedata.Comment{
		ID:        comment.ID,
		ArticleID: comment.ArticleID,
		Content:   comment.Content,
		User: templatedata.User{
			ID:       comment.UserID,
			Name:     comment.UserName,
			Nickname: comment.UserNickname,
		},
		ReplyUser: templatedata.User{
			ID:   comment.ReplyUserID,
			Name: comment.ReplyUserName,
		},
		Article: h.articleTemplateData(comment.Article),
	}
}

type articleActorHandler struct {
	userClientHandler
	contentClientHandler
	notifyUsecase *usecase.NotifyUsecase
}

type articleActorBuildReq struct {
	EventID   string
	ArticleID int64
	SenderID  int64
}

func (h *articleActorHandler) build(ctx context.Context, req *articleActorBuildReq) (*model.NotificationContext, error) {
	if req == nil {
		return nil, nil
	}
	var article *model.ContentArticle
	if h.contentClient != nil && req.ArticleID != 0 {
		articleResp, err := h.contentClient.GetArticle(ctx, req.ArticleID)
		if err != nil {
			return nil, err
		}
		article = articleResp
	}
	var authorID int64
	if article != nil {
		authorID = article.AuthorID
	}
	users, err := h.loadAccounts(ctx, req.SenderID, authorID)
	if err != nil {
		return nil, err
	}
	if article != nil {
		if author := users[article.AuthorID]; author != nil {
			article.AuthorName = author.Name
			article.AuthorNickname = author.Nickname
		}
	}
	templateData := templatedata.ArticleActor{
		Article: h.articleTemplateData(article),
		Actor:   h.templateUser(req.SenderID, users[req.SenderID]),
	}

	recipients := make([]*model.NotificationRecipient, 0, 1)
	if article != nil && article.AuthorID != 0 && article.AuthorID != req.SenderID {
		recipients = append(recipients, &model.NotificationRecipient{
			UserID: article.AuthorID,
		})
	}
	return &model.NotificationContext{
		EventID:      req.EventID,
		TemplateData: templateData,
		Recipients:   recipients,
	}, nil
}
