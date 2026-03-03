package util

import (
	"common/pkg/constant"
	"common/pkg/model"

	"context"
	"errors"
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
	SetCreatedBy(int64)
	SetUpdatedBy(int64)
	CreatedBy() (int64, bool)
	UpdatedBy() (int64, bool)
}

func UserAuditFields() []ent.Field {
	return []ent.Field{
		field.Int64("created_by").Comment("创建人ID").Nillable().Optional(),
		field.Int64("updated_by").Comment("更新人ID").Nillable().Optional(),
	}
}

type UsernameAuditSetter interface {
	SetCreatedByName(string)
	SetUpdatedByName(string)
	CreatedByName() (string, bool)
	UpdatedByName() (string, bool)
}

func UsernameAuditFields() []ent.Field {
	// 可选冗余字段，用于减少查询
	return []ent.Field{
		field.String("created_by_name").Comment("创建人用户名").Nillable().Optional(),
		field.String("updated_by_name").Comment("更新人用户名").Nillable().Optional(),
	}
}

func AuditHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			now := time.Now()
			user, userOk := GetContextValue[*model.User](ctx, constant.CtxUserInfo)
			var (
				userID   int64
				username string
			)
			if userOk {
				userID = user.ID
				username = user.Name
			}
			switch {
			case m.Op().Is(ent.OpCreate):
				// 时间字段
				if ts, ok := m.(TimeAuditSetter); ok {
					ts.SetCreatedAt(now)
					ts.SetUpdatedAt(now)
				}

				// 用户ID字段
				if us, ok := m.(UserAuditSetter); ok {
					if !userOk {
						return nil, errors.New("cannot get user info from context for UserAudit fields")
					}
					us.SetCreatedBy(userID)
					us.SetUpdatedBy(userID)
				}

				// 用户名字段
				if uns, ok := m.(UsernameAuditSetter); ok {
					if !userOk {
						return nil, errors.New("cannot get user info from context for UsernameAudit fields")
					}
					uns.SetCreatedByName(username)
					uns.SetUpdatedByName(username)
				}

			case m.Op().Is(ent.OpUpdate | ent.OpUpdateOne):
				// 时间字段
				if ts, ok := m.(TimeAuditSetter); ok {
					ts.SetUpdatedAt(now)
				}

				// 用户ID字段
				if us, ok := m.(UserAuditSetter); ok {
					if !userOk {
						return nil, errors.New("cannot get user info from context for UserAudit fields")
					}
					us.SetUpdatedBy(userID)
				}

				// 用户名字段
				if uns, ok := m.(UsernameAuditSetter); ok {
					if !userOk {
						return nil, errors.New("cannot get user info from context for UsernameAudit fields")
					}
					uns.SetUpdatedByName(username)
				}

			}
			return next.Mutate(ctx, m)
		})
	}
}
