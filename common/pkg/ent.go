package pkg

import (
	"context"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type TimeAuditSetter interface {
	SetCreatedAt(time.Time)
	SetUpdatedAt(time.Time)
	CreatedAt() (time.Time, bool)
	UpdatedAt() (time.Time, bool)
}

func TimeAuditFields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Comment("创建时间").Default(time.Now).Nillable(),
		field.Time("updated_at").Comment("更新时间").Default(time.Now).Nillable(),
	}
}

type UserAuditSetter interface {
	SetCreateBy(int)
	SetUpdateBy(int)
	CreateBy() (int, bool)
	UpdateBy() (int, bool)
}

func UserAuditFields() []ent.Field {
	return []ent.Field{
		field.Int("create_by").Comment("创建人ID").Optional(),
		field.Int("update_by").Comment("更新人ID").Optional(),
	}
}

var ContextUserIDKey = "user_id"

func ContextWithUserID(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, ContextUserIDKey, userID)
}

func ContextUserID(ctx context.Context) (int, bool) {
	v, ok := ctx.Value(ContextUserIDKey).(int)
	return v, ok
}

func AuditHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			now := time.Now()
			userID, _ := ContextUserID(ctx)

			switch {
			case m.Op().Is(ent.OpCreate):
				// 设置 created_at / updated_at
				if setter, ok := m.(interface{ SetCreatedAt(time.Time) }); ok {
					setter.SetCreatedAt(now)
				}
				if setter, ok := m.(interface{ SetUpdatedAt(time.Time) }); ok {
					setter.SetUpdatedAt(now)
				}
				// 设置 created_by / updated_by
				if setter, ok := m.(interface{ SetCreatedBy(int) }); ok && userID != 0 {
					setter.SetCreatedBy(userID)
				}
				if setter, ok := m.(interface{ SetUpdatedBy(int) }); ok && userID != 0 {
					setter.SetUpdatedBy(userID)
				}

			case m.Op().Is(ent.OpUpdate | ent.OpUpdateOne):
				// 只更新 updated_at / updated_by
				if setter, ok := m.(interface{ SetUpdatedAt(time.Time) }); ok {
					setter.SetUpdatedAt(now)
				}
				if setter, ok := m.(interface{ SetUpdatedBy(int) }); ok && userID != 0 {
					setter.SetUpdatedBy(userID)
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
