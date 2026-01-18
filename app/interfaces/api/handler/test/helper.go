package test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// PerformRequest 发起请求并断言状态码，返回解析后的响应体（泛型）
// Resp 应该是具体的响应结构体，例如 response.EditArticleResp 或 Result[T]
func PerformRequest[Resp any](t *testing.T, router *gin.Engine, method, path string, reqBody interface{}, expectedCode int) *Resp {
	t.Helper() // 标记为 helper 函数，报错时会跳过此函数的栈帧

	reqBytes, err := json.Marshal(reqBody)
	require.NoError(t, err, "failed to marshal request body")

	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewBuffer(reqBytes))
	req.Header.Set("Content-Type", "application/json")
	// 如果需要 token，可以在这里扩展，或者让 reqBody/Option 传入

	router.ServeHTTP(w, req)

	// 断言状态码
	assert.Equal(t, expectedCode, w.Code, "status code mismatch. Body: %s", w.Body.String())

	if w.Body.Len() == 0 {
		return nil
	}

	var resp Resp
	err = json.Unmarshal(w.Body.Bytes(), &resp)
	require.NoError(t, err, "failed to unmarshal response body: %s", w.Body.String())

	return &resp
}
