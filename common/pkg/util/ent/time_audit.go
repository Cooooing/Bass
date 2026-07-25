package ent

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

const (
	timeAuditCreatedAtFieldName = "created_at"
	timeAuditUpdatedAtFieldName = "updated_at"
)

type TimeAuditMixin struct{}

func (TimeAuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(timeAuditCreatedAtFieldName).Comment("创建时间").Default(time.Now).Nillable().Optional(),
		field.Time(timeAuditUpdatedAtFieldName).Comment("更新时间").Default(time.Now).Nillable().Optional(),
	}
}

func (TimeAuditMixin) Edges() []ent.Edge {
	return nil
}

func (TimeAuditMixin) Indexes() []ent.Index {
	return nil
}

func (TimeAuditMixin) Interceptors() []ent.Interceptor {
	return nil
}

func (TimeAuditMixin) Policy() ent.Policy {
	return nil
}

func (TimeAuditMixin) Annotations() []schema.Annotation {
	return nil
}

func (TimeAuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{timeAuditHook()}
}

type TimeAuditSetter interface {
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	CreatedAt() (time.Time, bool)
	UpdatedAt() (time.Time, bool)
}

func timeAuditHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			ts, ok := m.(TimeAuditSetter)
			if !ok {
				return next.Mutate(ctx, m)
			}
			now := time.Now()
			switch {
			case m.Op().Is(ent.OpCreate):
				ts.SetCreatedAt(now)
				ts.SetUpdatedAt(now)
			case m.Op().Is(ent.OpUpdate | ent.OpUpdateOne):
				ts.SetUpdatedAt(now)
			}
			return next.Mutate(ctx, m)
		})
	}
}
