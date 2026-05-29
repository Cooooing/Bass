package ent

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// --- 设置器接口 ---

type TimeAuditSetter interface {
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	CreatedAt() (time.Time, bool)
	UpdatedAt() (time.Time, bool)
}

type UserAuditSetter interface {
	SetCreatedBy(int64)
	SetUpdatedBy(int64)
	CreatedBy() (int64, bool)
	UpdatedBy() (int64, bool)
}

// --- TimeAuditMixin ---

type TimeAuditMixin struct{}

func (TimeAuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间").Nillable().Optional(),
		field.Time("updated_at").Comment("更新时间").Nillable().Optional(),
	}
}

func (TimeAuditMixin) Edges() []ent.Edge                { return nil }
func (TimeAuditMixin) Indexes() []ent.Index             { return nil }
func (TimeAuditMixin) Interceptors() []ent.Interceptor  { return nil }
func (TimeAuditMixin) Policy() ent.Policy               { return nil }
func (TimeAuditMixin) Annotations() []schema.Annotation { return nil }

func (TimeAuditMixin) Hooks() []ent.Hook {
	return []ent.Hook{timeAuditHook()}
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

// --- UserAuditMixin ---

type UserAuditMixin struct{}

func (UserAuditMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
}

func (UserAuditMixin) Edges() []ent.Edge                { return nil }
func (UserAuditMixin) Indexes() []ent.Index             { return nil }
func (UserAuditMixin) Interceptors() []ent.Interceptor  { return nil }
func (UserAuditMixin) Policy() ent.Policy               { return nil }
func (UserAuditMixin) Annotations() []schema.Annotation { return nil }

func (UserAuditMixin) Hooks() []ent.Hook {
	return nil
}
