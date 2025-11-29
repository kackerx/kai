package handler

import (
	"context"

	"github.com/kackerx/kai/app/interfaces/api/request"
	"github.com/kackerx/kai/app/interfaces/api/response"
)

type Test interface {
	TestWrap(ctx context.Context, req *request.TestWrapReq) (*response.TestWrapResp, error)
}

func NewTest() Test {
	return &testHandler{}
}

type testHandler struct {
}

func (a *testHandler) TestWrap(ctx context.Context, req *request.TestWrapReq) (*response.TestWrapResp, error) {

	return &response.TestWrapResp{
		Message: "hello " + req.Name,
	}, nil
}
