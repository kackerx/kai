package application

import (
	"context"
	"fmt"

	"github.com/kackerx/kai/app/domain/article"
	"github.com/kackerx/kai/app/interfaces/api/request"
)

type Article interface {
	Edit(ctx context.Context, req *request.EditArticleReq) (int64, error)
}

func NewArticle(articleRepo article.Repository) Article {
	return &articleApp{
		repo: articleRepo,
	}
}

type articleApp struct {
	repo article.Repository
}

func (a *articleApp) Edit(ctx context.Context, req *request.EditArticleReq) (int64, error) {
	return 1, nil
}

type Person struct {
	Name string
}

func (p *Person) SayHello() {
	fmt.Println("Hello, my name is", p.Name)
}
