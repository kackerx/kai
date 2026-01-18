package comment

import (
	"time"

	"github.com/kackerx/kai/app/domain/comment"
	"github.com/kackerx/kai/pkg/util/structure"
)

type Model struct {
	ID        string    `gorm:"column:id;primary_key;size:36;"`
	UserID    string    `gorm:"column:user_id;size:36;index;not null;"`
	BizType   string    `gorm:"column:biz_type;size:50;index;not null;"`
	BizID     string    `gorm:"column:biz_id;size:36;index;not null;"`
	ParentID  string    `gorm:"column:parent_id;size:36;index;default:'';not null;"`
	RootID    string    `gorm:"column:root_id;size:36;index;default:'';not null;"`
	Content   string    `gorm:"column:content;type:text;not null;"`
	CreatedAt time.Time `gorm:"column:created_at;index;"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;"`
}

func (Model) TableName() string {
	return "comments"
}

func (a Model) ToDomain() *comment.Comment {
	item := new(comment.Comment)
	structure.Copy(a, item)
	return item
}

func toDomainList(ms []*Model) comment.Comments {
	list := make(comment.Comments, len(ms))
	for i, item := range ms {
		list[i] = item.ToDomain()
	}
	return list
}

func domainToModel(c *comment.Comment) *Model {
	item := new(Model)
	structure.Copy(c, item)
	return item
}
