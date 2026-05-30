package data

import (
	"bbs/internal/biz/repo"
	bbscontentv1 "common/api/gen/bbs/v1/content"
	bbsuserv1 "common/api/gen/bbs/v1/user"
	contentv1 "common/api/gen/content/v1"
	"common/pkg/client/rpc"
	"context"
)

var _ repo.ContentCommentRepo = (*ContentCommentRepo)(nil)

type ContentCommentRepo struct {
	contentClient *rpc.ContentClient
}

func NewContentCommentRepo(contentClient *rpc.ContentClient) repo.ContentCommentRepo {
	return &ContentCommentRepo{contentClient: contentClient}
}

func (r *ContentCommentRepo) CreateComment(ctx context.Context, req *bbscontentv1.CreateComment_Request) (*bbscontentv1.CreateComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	reply, err := r.contentClient.Comment.Create(ctx, &contentv1.CreateComment_Request{ArticleId: req.GetArticleId(), Content: req.GetContent(), ReplyId: req.GetReplyId(), UserId: userID})
	if err != nil {
		return nil, err
	}
	item := reply.GetComment()
	return &bbscontentv1.CreateComment_Reply{Comment: &bbscontentv1.Comment{
		Id:            item.GetId(),
		ArticleId:     item.GetArticleId(),
		Content:       item.GetContent(),
		ContentRender: item.GetContentRender(),
		Level:         item.GetLevel(),
		ParentId:      item.ParentId,
		ReplyId:       item.ReplyId,
		Status:        bbscontentv1.CommentStatus(item.GetStatus()),
		ThankCount:    item.GetThankCount(),
		LikeCount:     item.GetLikeCount(),
		ReplyCount:    item.GetReplyCount(),
		CreatedBy:     item.CreatedBy,
		UpdatedBy:     item.UpdatedBy,
		CreatedAt:     formatProtoTime(item.GetCreatedAt()),
		UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
	}}, nil
}

func (r *ContentCommentRepo) ListComments(ctx context.Context, req *bbscontentv1.ListComments_Request) (*bbscontentv1.ListComments_Reply, error) {
	query := req.GetQuery()
	if query == nil {
		query = &bbscontentv1.CommentQuery{}
	}
	contentQuery := &contentv1.CommentQueryParams{
		CommentId: query.CommentId,
		ArticleId: query.ArticleId,
		ParentId:  query.ParentId,
		ReplyId:   query.ReplyId,
		UserId:    query.UserId,
		Level:     query.Level,
	}
	if query.Order != nil {
		contentQuery.Order = new(contentv1.CommentOrder(*query.Order))
	}
	if query.Status != nil {
		contentQuery.Status = new(contentv1.CommentStatus(*query.Status))
	}
	reply, err := r.contentClient.Comment.List(ctx, &contentv1.ListComments_Request{
		Page:  req.Page,
		Query: contentQuery,
	})
	if err != nil {
		return nil, err
	}
	rows := make([]*bbscontentv1.Comment, 0, len(reply.GetRows()))
	for _, item := range reply.GetRows() {
		row := &bbscontentv1.Comment{
			Id:            item.GetId(),
			ArticleId:     item.GetArticleId(),
			Content:       item.GetContent(),
			ContentRender: item.GetContentRender(),
			Level:         item.GetLevel(),
			ParentId:      item.ParentId,
			ReplyId:       item.ReplyId,
			Status:        bbscontentv1.CommentStatus(item.GetStatus()),
			ThankCount:    item.GetThankCount(),
			LikeCount:     item.GetLikeCount(),
			ReplyCount:    item.GetReplyCount(),
			CreatedBy:     item.CreatedBy,
			UpdatedBy:     item.UpdatedBy,
			CreatedAt:     formatProtoTime(item.GetCreatedAt()),
			UpdatedAt:     formatProtoTime(item.GetUpdatedAt()),
		}
		if user := item.GetUser(); user != nil {
			row.User = &bbsuserv1.AccountProfile{Id: user.GetId(), Name: user.GetName(), Nickname: user.Nickname, AvatarUrl: user.AvatarUrl}
		}
		if replyUser := item.GetReplyUser(); replyUser != nil {
			row.ReplyUser = &bbsuserv1.AccountProfile{Id: replyUser.GetId(), Name: replyUser.GetName(), Nickname: replyUser.Nickname, AvatarUrl: replyUser.AvatarUrl}
		}
		rows = append(rows, row)
	}
	return &bbscontentv1.ListComments_Reply{Page: reply.GetPage(), Rows: rows}, nil
}

func (r *ContentCommentRepo) LikeComment(ctx context.Context, req *bbscontentv1.LikeComment_Request) (*bbscontentv1.LikeComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Comment.Like(ctx, &contentv1.LikeComment_Request{Id: req.GetId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.LikeComment_Reply{}, nil
}

func (r *ContentCommentRepo) ThankComment(ctx context.Context, req *bbscontentv1.ThankComment_Request) (*bbscontentv1.ThankComment_Reply, error) {
	userID, err := currentUserID(ctx)
	if err != nil {
		return nil, err
	}
	_, err = r.contentClient.Comment.Thank(ctx, &contentv1.ThankComment_Request{Id: req.GetId(), Active: req.GetActive(), UserId: userID})
	if err != nil {
		return nil, err
	}
	return &bbscontentv1.ThankComment_Reply{}, nil
}
