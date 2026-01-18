package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/interfaces/api/request"
	"github.com/kackerx/kai/app/interfaces/api/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

// ArticleTestSuite 测试套件
type ArticleTestSuite struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
}

func (s *ArticleTestSuite) SetupSuite() {
	engine, db, _, err := BuildEngine()
	require.NoError(s.T(), err, "failed to build engine")

	s.server = engine
	s.db = db
}

func (s *ArticleTestSuite) TearDownTest() {
	fmt.Println("TRUNCATE TABLE articlesj")
}

func (s *ArticleTestSuite) TestEdit() {
	t := s.T()

	testCase := []struct {
		name string

		before func(t *testing.T)
		after  func(t *testing.T)

		req *request.EditArticleReq

		wantCode   int
		wantResult Result[int64]
	}{
		{
			name: "保存成功",
			before: func(t *testing.T) {
				// TODO: 准备数据，例如插入一条文章记录以便更新，或者保持清空以便新建
			},
			after: func(t *testing.T) {
				// TODO: 清理数据
			},
			req: &request.EditArticleReq{
				Title:   "test title",
				Content: "test content",
			},
			wantCode: 200,
			wantResult: Result[int64]{
				Code:    0,
				Message: "success",
				Data:    1,
			},
		},
	}

	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			if tt.before != nil {
				tt.before(t)
			}

			body, err := json.Marshal(tt.req)
			require.NoError(err, "failed to marshal req")

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/articles", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			s.server.ServeHTTP(w, req)

			assert.Equal(tt.wantCode, w.Code)

			var got response.CommonResult
			err = json.Unmarshal(w.Body.Bytes(), &got)
			require.NoError(err)

			assert.Equal(tt.wantResult.Code, got.Code)

			if tt.after != nil {
				tt.after(t)
			}
		})
	}

}

func (s *ArticleTestSuite) TestAssert() {
	t := s.T()
	assert := assert.New(t)

	a := &Result[int64]{
		Code:    1,
		Message: "success",
		Data:    1,
	}
	b := &Result[int64]{
		Code:    1,
		Message: "success",
		Data:    1,
	}

	assert.Equal(a, b)

	assert.Equal(errors.New("test error"), "test eror")
	c, d := errors.New("test error"), errors.New("test rror")
	assert.Equal(c, d)
}

type Result[T any] struct {
	Code    int
	Message string
	Data    T
}

func TestArticle(t *testing.T) {
	suite.Run(t, new(ArticleTestSuite))
}
