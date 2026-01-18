package article

import (
	"github.com/kackerx/kai/app/domain/article"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) article.Repository {
	return &repository{
		db: db,
	}
}

type repository struct {
	db *gorm.DB
}
