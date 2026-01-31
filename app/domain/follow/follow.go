package follow

import "time"

type FollowRelation struct {
	ID int

	// 联合索引
	Followee int
	Follower int

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
