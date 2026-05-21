package event

type OutboxStatus string

const (
	OutboxStatusPending    OutboxStatus = "pending"
	OutboxStatusPublishing OutboxStatus = "publishing"
	OutboxStatusPublished  OutboxStatus = "published"
	OutboxStatusFailed     OutboxStatus = "failed"
	OutboxStatusDead       OutboxStatus = "dead"
)

func OutboxStatusValues() []string {
	return []string{
		string(OutboxStatusPending),
		string(OutboxStatusPublishing),
		string(OutboxStatusPublished),
		string(OutboxStatusFailed),
		string(OutboxStatusDead),
	}
}

type InboxStatus string

const (
	InboxStatusReceived   InboxStatus = "received"
	InboxStatusProcessing InboxStatus = "processing"
	InboxStatusProcessed  InboxStatus = "processed"
	InboxStatusFailed     InboxStatus = "failed"
	InboxStatusDead       InboxStatus = "dead"
)

func InboxStatusValues() []string {
	return []string{
		string(InboxStatusReceived),
		string(InboxStatusProcessing),
		string(InboxStatusProcessed),
		string(InboxStatusFailed),
		string(InboxStatusDead),
	}
}
