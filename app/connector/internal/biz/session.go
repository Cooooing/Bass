package biz

import "common/pkg/cutil/collections/set"

type SessionDomain struct {
	*BaseDomain
	sessionIds set.Set[string]
}

func NewSessionDomain(baseDomain *BaseDomain) *SessionDomain {
	return &SessionDomain{
		BaseDomain: baseDomain,
		sessionIds: set.NewThreadSafeComparableSet[string](0),
	}
}

func (d *SessionDomain) AddSessionId(sessionId string) {
	d.sessionIds.Add(sessionId)
}

func (d *SessionDomain) RemoveSessionId(sessionId string) {
	d.sessionIds.Remove(sessionId)
}

func (d *SessionDomain) GetSessionIds() []string {
	return d.sessionIds.ToSlice()
}
