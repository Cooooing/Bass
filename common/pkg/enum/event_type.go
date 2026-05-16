package enum

// EventType 事件类型，字符串值与 proto enum 名称完全一致
type EventType string

const (
	EventTypeUserFollowCreated EventType = "EVENT_TYPE_USER_FOLLOW_CREATED"
	EventTypeUserFollowDeleted EventType = "EVENT_TYPE_USER_FOLLOW_DELETED"
	EventTypeArticlePublished  EventType = "EVENT_TYPE_ARTICLE_PUBLISHED"
	EventTypeArticleLiked      EventType = "EVENT_TYPE_ARTICLE_LIKED"
	EventTypeArticleThanked    EventType = "EVENT_TYPE_ARTICLE_THANKED"
	EventTypeArticleCollected  EventType = "EVENT_TYPE_ARTICLE_COLLECTED"
	EventTypeArticleWatched    EventType = "EVENT_TYPE_ARTICLE_WATCHED"
	EventTypeCommentPublished  EventType = "EVENT_TYPE_COMMENT_PUBLISHED"
	EventTypeCommentLiked      EventType = "EVENT_TYPE_COMMENT_LIKED"
)

func (EventType) Values() []string {
	return []string{
		string(EventTypeUserFollowCreated),
		string(EventTypeUserFollowDeleted),
		string(EventTypeArticlePublished),
		string(EventTypeArticleLiked),
		string(EventTypeArticleThanked),
		string(EventTypeArticleCollected),
		string(EventTypeArticleWatched),
		string(EventTypeCommentPublished),
		string(EventTypeCommentLiked),
	}
}
