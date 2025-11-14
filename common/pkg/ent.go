package pkg

import (
	"common/pkg/util"
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
		field.Time("created_at").Comment("创建时间").Default(time.Now).Nillable().Optional(),
		field.Time("updated_at").Comment("更新时间").Default(time.Now).Nillable().Optional(),
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
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
}

func AuditHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			now := time.Now()
			user, userIdOk := util.GetUserInfo(ctx)
			var userID int64
			if userIdOk {
				userID = user.ID
			}
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
				if setter, ok := m.(interface{ SetCreatedBy(int64) }); ok && userID != 0 {
					if !userIdOk {
						panic("can not get userId")
					}
					setter.SetCreatedBy(userID)
				}
				if setter, ok := m.(interface{ SetUpdatedBy(int64) }); ok && userID != 0 {
					if userIdOk {
						setter.SetUpdatedBy(userID)
					}
				}

			case m.Op().Is(ent.OpUpdate | ent.OpUpdateOne):
				// 只更新 updated_at / updated_by
				if setter, ok := m.(interface{ SetUpdatedAt(time.Time) }); ok {
					setter.SetUpdatedAt(now)
				}
				if setter, ok := m.(interface{ SetUpdatedBy(int64) }); ok && userID != 0 {
					if userIdOk {
						setter.SetUpdatedBy(userID)
					}
				}
			}
			return next.Mutate(ctx, m)
		})
	}
}
