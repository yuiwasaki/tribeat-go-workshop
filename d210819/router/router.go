package router

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

type APIHandler struct{}

// NewAPIHandler APIHandlerを作成する
func NewAPIHandler() APIHandler {
	return APIHandler{}
}

func ErrorResult(ctx echo.Context, status int, code, detail string) error {
	result := oapi.ResultError{
		Result: false,
		Error: oapi.Error{
			Detail:     detail,
			ErrorCode:  code,
			StatusCode: float32(status),
		},
	}
	return ctx.JSON(status, result)
}

// グループ一覧取得
// (GET /api/client/groups)
func (APIHandler) GetApiClientGroups(ctx echo.Context) error {
	return nil
}

// グループ作成
// (POST /api/client/groups)
func (APIHandler) PostApiClientGroups(ctx echo.Context) error {
	return nil
}

// ユーザー一覧取得
// (GET /api/client/users)
func (APIHandler) GetApiClientUsers(ctx echo.Context) error {
	return nil
}

// ユーザー新規登録
// (POST /api/client/users)
func (APIHandler) PostApiClientUsers(ctx echo.Context) error {
	return nil
}

// ユーザー一覧取得
// (GET /api/client/users/{userId})
func (APIHandler) GetApiClientUsersUserId(ctx echo.Context, userId oapi.UserId) error {
	return nil
}

// ヘルスチェック
// (GET /api/healthcheck)
func (APIHandler) GetApiHealthcheck(ctx echo.Context) error {
	return nil
}

// ログイン
// (POST /api/login)
func (APIHandler) PostApiLogin(ctx echo.Context) error {
	var req oapi.RequestLogin
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return ErrorResult(ctx, http.StatusInternalServerError, "ERROR", err.Error())
	}
	return ctx.JSON(http.StatusCreated, nil)
}

// RegisterHandler ハンドラーをセットする
func (handler APIHandler) RegisterHandler(router *echo.Echo) {
	wrapper := oapi.ServerInterfaceWrapper{
		Handler: handler,
	}
	router.GET("/api/healthcheck", wrapper.GetApiHealthcheck)
}
