package comment

import (
	"context"

	"github.com/kackerx/kai/app/domain/comment"
	"github.com/kackerx/kai/app/domain/errors"
	"github.com/kackerx/kai/app/domain/pagination"
	"github.com/kackerx/kai/app/infrastructure/gormx"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) comment.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}

func GetModelDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return gormx.GetDBWithModel(ctx, defDB, new(Model))
}

func (a *repository) Get(ctx context.Context, id string) (*comment.Comment, error) {
	item := &Model{}
	ok, err := gormx.FindOne(ctx, GetModelDB(ctx, a.db).Where("id=?", id), item)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !ok {
		return nil, nil
	}
	return item.ToDomain(), nil
}

func (a *repository) List(ctx context.Context, params comment.QueryParams) (comment.Comments, *pagination.Pagination, error) {
	db := GetModelDB(ctx, a.db)
	if v := params.BizType; v != "" {
		db = db.Where("biz_type=?", v)
	}
	if v := params.BizID; v != "" {
		db = db.Where("biz_id=?", v)
	}
	if v := params.ParentID; v != "" {
		db = db.Where("parent_id=?", v)
	}
	if v := params.RootID; v != "" {
		db = db.Where("root_id=?", v)
	}
	if v := params.UserID; v != "" {
		db = db.Where("user_id=?", v)
	}

	db = db.Order(gormx.ParseOrder(params.OrderFields.AddIdSortField()))

	var list []*Model
	pr, err := gormx.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, nil, errors.WithStack(err)
	}
	return toDomainList(list), pr, nil
}

func (a *repository) Create(ctx context.Context, c *comment.Comment) error {
	result := GetModelDB(ctx, a.db).Create(domainToModel(c))
	return errors.WithStack(result.Error)
}

func (a *repository) Update(ctx context.Context, c *comment.Comment) error {
	result := GetModelDB(ctx, a.db).Where("id=?", c.ID).Updates(domainToModel(c))
	return errors.WithStack(result.Error)
}

func (a *repository) Delete(ctx context.Context, id string) error {
	result := GetModelDB(ctx, a.db).Where("id=?", id).Delete(Model{})
	return errors.WithStack(result.Error)
}
