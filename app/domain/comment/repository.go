package comment

import (
	"context"

	"github.com/kackerx/kai/app/domain/pagination"
)

type Repository interface {
	Create(ctx context.Context, comment *Comment) error
	Update(ctx context.Context, comment *Comment) error
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*Comment, error)
	List(ctx context.Context, params QueryParams) (Comments, *pagination.Pagination, error)
}

type QueryParams struct {
	PaginationParam pagination.Param
	OrderFields     pagination.OrderFields
	BizType         BizType
	BizID           string
	ParentID        string
	RootID          string
	UserID          string
}
