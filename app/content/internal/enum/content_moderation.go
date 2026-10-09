package enum

import (
	"common/pkg/enum"
	v1 "common/proto/gen/content/v1/enum"
)

// ContentModerationTarget 表示内容管理处置目标类型。
type ContentModerationTarget string

const (
	ContentModerationTargetArticle    ContentModerationTarget = "article"    // ContentModerationTargetArticle 表示处置目标为文章。
	ContentModerationTargetComment    ContentModerationTarget = "comment"    // ContentModerationTargetComment 表示处置目标为评论。
	ContentModerationTargetPostscript ContentModerationTarget = "postscript" // ContentModerationTargetPostscript 表示处置目标为附言。
)

// ContentModerationTargetMap 维护内容管理目标内部枚举与 proto 枚举的映射。
var ContentModerationTargetMap = enum.NewMapping[ContentModerationTarget, v1.ContentModerationTarget](map[ContentModerationTarget]enum.Entry[ContentModerationTarget, v1.ContentModerationTarget]{
	ContentModerationTargetArticle:    {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_ARTICLE},
	ContentModerationTargetComment:    {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_COMMENT},
	ContentModerationTargetPostscript: {Proto: v1.ContentModerationTarget_CONTENT_MODERATION_TARGET_POSTSCRIPT},
})

func (e ContentModerationTarget) String() string {
	return string(e)
}
