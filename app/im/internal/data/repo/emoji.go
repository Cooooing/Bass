package repo

import (
	"context"
	"fmt"

	utilent "common/pkg/util/ent"

	"im/internal/biz/model"
	bizrepo "im/internal/biz/repo"
	"im/internal/data/gen"
	"im/internal/data/gen/emojigroup"
	"im/internal/data/gen/emojiitem"
)

// EmojiGroupRepo 将表情分组领域模型映射到 IM 数据库。
type EmojiGroupRepo struct{ db *gen.Client }

func NewEmojiGroupRepo(db *gen.Client) bizrepo.EmojiGroupRepo { return &EmojiGroupRepo{db: db} }

func (r *EmojiGroupRepo) client(ctx context.Context) *gen.Client {
	if tx, ok := utilent.ClientFromCtx[*gen.Client](ctx); ok {
		return tx
	}
	return r.db
}

func (r *EmojiGroupRepo) List(ctx context.Context, userID int64, query string) ([]*model.EmojiGroup, error) {
	rows, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.UserIDEQ(userID)).Order(gen.Desc(emojigroup.FieldSort), gen.Desc(emojigroup.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	groups := make([]*model.EmojiGroup, 0, len(rows))
	byID := make(map[int64]*model.EmojiGroup, len(rows))
	ids := make([]int64, 0, len(rows))
	for _, row := range rows {
		group := r.group(row)
		groups = append(groups, group)
		byID[group.ID] = group
		ids = append(ids, group.ID)
	}
	if len(ids) == 0 {
		return groups, nil
	}
	itemQuery := r.client(ctx).EmojiItem.Query().Where(emojiitem.GroupIDIn(ids...))
	if query != "" {
		itemQuery = itemQuery.Where(emojiitem.NameContainsFold(query))
	}
	items, err := itemQuery.Order(gen.Desc(emojiitem.FieldSort), gen.Desc(emojiitem.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, err
	}
	for _, item := range items {
		if item.GroupID != nil {
			if group := byID[*item.GroupID]; group != nil {
				group.Items = append(group.Items, r.item(item))
			}
		}
	}
	if query == "" {
		return groups, nil
	}
	matched := make([]*model.EmojiGroup, 0, len(groups))
	for _, group := range groups {
		if len(group.Items) > 0 {
			matched = append(matched, group)
		}
	}
	return matched, nil
}

func (r *EmojiGroupRepo) Get(ctx context.Context, userID, groupID int64) (*model.EmojiGroup, error) {
	row, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.IDEQ(groupID), emojigroup.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.group(row), nil
}

func (r *EmojiGroupRepo) GetDefault(ctx context.Context, userID int64) (*model.EmojiGroup, error) {
	row, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.UserIDEQ(userID), emojigroup.NameEQ("默认分组")).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return r.group(row), nil
}

func (r *EmojiGroupRepo) Create(ctx context.Context, group *model.EmojiGroup) (*model.EmojiGroup, error) {
	row, err := r.client(ctx).EmojiGroup.Create().SetUserID(group.UserID).SetName(group.Name).SetSort(group.Sort).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.group(row), nil
}

func (r *EmojiGroupRepo) Rename(ctx context.Context, userID, groupID int64, name string) (*model.EmojiGroup, error) {
	row, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.IDEQ(groupID), emojigroup.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	row, err = r.client(ctx).EmojiGroup.UpdateOneID(row.ID).SetName(name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.group(row), nil
}

func (r *EmojiGroupRepo) Delete(ctx context.Context, userID, groupID int64) error {
	client := r.client(ctx)
	if _, err := client.EmojiItem.Delete().Where(emojiitem.GroupIDEQ(groupID), emojiitem.UserIDEQ(userID)).Exec(ctx); err != nil {
		return err
	}
	_, err := client.EmojiGroup.Delete().Where(emojigroup.IDEQ(groupID), emojigroup.UserIDEQ(userID)).Exec(ctx)
	return err
}

func (r *EmojiGroupRepo) CreateItem(ctx context.Context, item *model.EmojiItem) (*model.EmojiItem, error) {
	creator := r.client(ctx).EmojiItem.Create().SetUserID(item.UserID).SetAssetID(item.AssetID).SetSort(item.Sort)
	if item.GroupID != nil {
		creator.SetGroupID(*item.GroupID)
	}
	if item.Name != "" {
		creator.SetName(item.Name)
	}
	row, err := creator.Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.item(row), nil
}

func (r *EmojiGroupRepo) RenameItem(ctx context.Context, userID, itemID int64, name string) (*model.EmojiItem, error) {
	row, err := r.client(ctx).EmojiItem.Query().Where(emojiitem.IDEQ(itemID), emojiitem.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	row, err = r.client(ctx).EmojiItem.UpdateOneID(row.ID).SetName(name).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.item(row), nil
}

func (r *EmojiGroupRepo) DeleteItem(ctx context.Context, userID, itemID int64) error {
	_, err := r.client(ctx).EmojiItem.Delete().Where(emojiitem.IDEQ(itemID), emojiitem.UserIDEQ(userID)).Exec(ctx)
	return err
}

func (r *EmojiGroupRepo) MoveItem(ctx context.Context, userID, itemID, groupID int64) (*model.EmojiItem, error) {
	client := r.client(ctx)
	group, err := client.EmojiGroup.Query().Where(emojigroup.IDEQ(groupID), emojigroup.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	item, err := client.EmojiItem.Query().Where(emojiitem.IDEQ(itemID), emojiitem.UserIDEQ(userID)).Only(ctx)
	if gen.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	if item.GroupID != nil && *item.GroupID == group.ID {
		return r.item(item), nil
	}
	item, err = client.EmojiItem.UpdateOneID(item.ID).SetGroupID(group.ID).SetSort(0).Save(ctx)
	if err != nil {
		return nil, err
	}
	return r.item(item), nil
}

func (r *EmojiGroupRepo) ReorderGroups(ctx context.Context, userID int64, ids []int64) error {
	rows, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.UserIDEQ(userID), emojigroup.IDIn(ids...)).All(ctx)
	if err != nil {
		return err
	}
	total, err := r.client(ctx).EmojiGroup.Query().Where(emojigroup.UserIDEQ(userID)).Count(ctx)
	if err != nil {
		return err
	}
	if len(rows) != len(ids) || total != len(ids) {
		return fmt.Errorf("emoji groups do not belong to user")
	}
	for index, id := range ids {
		if err = r.client(ctx).EmojiGroup.UpdateOneID(id).SetSort(int64(len(ids) - index)).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r *EmojiGroupRepo) ReorderItems(ctx context.Context, userID int64, groupID *int64, ids []int64) error {
	query := r.client(ctx).EmojiItem.Query().Where(emojiitem.UserIDEQ(userID), emojiitem.IDIn(ids...))
	allItemsQuery := r.client(ctx).EmojiItem.Query().Where(emojiitem.UserIDEQ(userID))
	if groupID == nil {
		query = query.Where(emojiitem.GroupIDIsNil())
		allItemsQuery = allItemsQuery.Where(emojiitem.GroupIDIsNil())
	} else {
		query = query.Where(emojiitem.GroupIDEQ(*groupID))
		allItemsQuery = allItemsQuery.Where(emojiitem.GroupIDEQ(*groupID))
	}
	rows, err := query.All(ctx)
	if err != nil {
		return err
	}
	total, err := allItemsQuery.Count(ctx)
	if err != nil {
		return err
	}
	if len(rows) != len(ids) || total != len(ids) {
		return fmt.Errorf("emoji items do not belong to group")
	}
	for index, id := range ids {
		if err = r.client(ctx).EmojiItem.UpdateOneID(id).SetSort(int64(len(ids) - index)).Exec(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (*EmojiGroupRepo) group(row *gen.EmojiGroup) *model.EmojiGroup {
	return &model.EmojiGroup{ID: row.ID, UserID: row.UserID, Name: row.Name, Sort: row.Sort, CreatedAt: row.CreatedAt, Items: make([]*model.EmojiItem, 0)}
}

func (*EmojiGroupRepo) item(row *gen.EmojiItem) *model.EmojiItem {
	return &model.EmojiItem{ID: row.ID, UserID: row.UserID, GroupID: row.GroupID, AssetID: row.AssetID, Name: row.Name, Sort: row.Sort}
}
