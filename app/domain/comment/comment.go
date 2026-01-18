package comment

import "time"

type Comment struct {
	ID        string
	UserID    string
	BizType   BizType
	BizID     string
	ParentID  string
	RootID    string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comments []*Comment
