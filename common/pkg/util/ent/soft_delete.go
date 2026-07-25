package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

const softDeleteFieldName = "deleted_at"

type SoftDeleteMixin struct{}

func (SoftDeleteMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time(softDeleteFieldName).Comment("逻辑删除时间").Nillable().Optional(),
	}
}

func (SoftDeleteMixin) Edges() []ent.Edge {
	return nil
}

func (SoftDeleteMixin) Indexes() []ent.Index {
	return nil
}

func (SoftDeleteMixin) Interceptors() []ent.Interceptor {
	return nil
}

func (SoftDeleteMixin) Policy() ent.Policy {
	return nil
}

func (SoftDeleteMixin) Annotations() []schema.Annotation {
	return nil
}

func (SoftDeleteMixin) Hooks() []ent.Hook {
	return []ent.Hook{softDeleteHook()}
}

type softDeleteMutation interface {
	softDeleteWhereMutation
	SetOp(ent.Op)
	SetDeletedAt(time.Time)
}

type softDeleteWhereMutation interface {
	WhereP(...func(*sql.Selector))
}

func softDeleteHook() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if m.Op().Is(ent.OpUpdate | ent.OpUpdateOne | ent.OpDelete | ent.OpDeleteOne) {
				mutation, ok := m.(softDeleteWhereMutation)
				if !ok {
					return nil, fmt.Errorf("soft delete unsupported mutation %T", m)
				}
				mutation.WhereP(sql.FieldIsNull(softDeleteFieldName))
			}
			if !m.Op().Is(ent.OpDelete | ent.OpDeleteOne) {
				return next.Mutate(ctx, m)
			}
			mutation, ok := m.(softDeleteMutation)
			if !ok {
				return nil, fmt.Errorf("soft delete unsupported mutation %T", m)
			}
			mutation.SetOp(ent.OpUpdate)
			mutation.SetDeletedAt(time.Now())
			return next.Mutate(ctx, m)
		})
	}
}
