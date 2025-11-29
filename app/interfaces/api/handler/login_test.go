package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	appmock "github.com/kackerx/kai/app/application/mock"
	"github.com/kackerx/kai/app/domain/user"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestLogin_Login(t *testing.T) {
	assert := assert.New(t)
	// 1. 创建一个 ResponseRecorder 来记录响应
	w := httptest.NewRecorder()

	// 2. 创建一个 Gin 的上下文，Gin 在测试模式下运行
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	app := appmock.NewMockLogin(ctrl)
	app.EXPECT().Verify(
		gomock.Any(),
		"admin",
		"123456",
	).Return(&user.User{
		ID:       "2",
		UserName: "admin",
		Password: "123456",
	}, nil)

	loginHandler := NewLogin(app)
	// 注册你的路由和处理器
	router.POST("/api/v1/pub/login", loginHandler.Login)

	// 3. 创建一个请求
	req, err := http.NewRequest(http.MethodPost, "/api/v1/pub/login", bytes.NewBuffer([]byte(`
	{
		"user_name": "admin",
		"password": "123456",
		"captcha_id": "admin",
		"captcha_code": "123456"
	}
	`)))
	assert.NoError(err)
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(http.StatusOK, w.Code)
}
