package model

type UserAccount struct {
	ID       int64
	Name     string
	Nickname string
	Phone    string
	Email    string
}

type ContentArticle struct {
	ID             int64
	Title          string
	AuthorID       int64
	AuthorName     string
	AuthorNickname string
}

type ContentComment struct {
	ID            int64
	ArticleID     int64
	Content       string
	UserID        int64
	UserName      string
	UserNickname  string
	ReplyUserID   int64
	ReplyUserName string
	Article       *ContentArticle
}
