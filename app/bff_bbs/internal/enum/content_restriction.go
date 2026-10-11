package enum

import (
	commonenum "common/pkg/enum"
	bffv1 "common/proto/gen/bff_bbs/v1/content/enum"
	contentv1 "common/proto/gen/content/v1/enum"
)

type ContentRestriction string

const (
	ContentRestrictionNone   ContentRestriction = "none"   // none
	ContentRestrictionHidden ContentRestriction = "hidden" // hidden
	ContentRestrictionLocked ContentRestriction = "locked" // locked
)

var ContentRestrictionMap = commonenum.NewMapping[ContentRestriction, bffv1.ContentRestriction](map[ContentRestriction]commonenum.Entry[ContentRestriction, bffv1.ContentRestriction]{
	ContentRestrictionNone:   {Proto: bffv1.ContentRestriction_CONTENT_RESTRICTION_NONE},
	ContentRestrictionHidden: {Proto: bffv1.ContentRestriction_CONTENT_RESTRICTION_HIDDEN},
	ContentRestrictionLocked: {Proto: bffv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED},
})

type ContentRestrictionConverter struct{}

func NewContentRestrictionConverter() *ContentRestrictionConverter {
	return &ContentRestrictionConverter{}
}

func (*ContentRestrictionConverter) FromContentProto(value contentv1.ContentRestriction) ContentRestriction {
	switch value {
	case contentv1.ContentRestriction_CONTENT_RESTRICTION_HIDDEN:
		return ContentRestrictionHidden
	case contentv1.ContentRestriction_CONTENT_RESTRICTION_LOCKED:
		return ContentRestrictionLocked
	default:
		return ContentRestrictionNone
	}
}
