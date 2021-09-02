package router

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

type APIHandler struct{}

type APIError struct {
	Code    int
	Message string
}

func NewAPIError(code int, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

func (err *APIError) Error() string {
	return fmt.Sprintf("Code=%d Message=%s", err.Code, err.Message)
}

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
	if req.Email != "hoge@hoge.com" {
		return fmt.Errorf("HOGE")
	}
	res := oapi.ResultOK{
		Result: true,
	}
	return ctx.JSON(http.StatusCreated, res)
}

func errorHandleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 処理
		err := next(ctx)
		if err != nil {
			if e, ok := err.(*echo.HTTPError); ok {
				return ErrorResult(ctx, e.Code, "BAD REQUEST", fmt.Sprintf("パラメーターが正しくない[%v]", e.Message))
			}
			if e, ok := err.(*APIError); ok {
				return ErrorResult(ctx, e.Code, "API ERROR", fmt.Sprintf("APIでエラーが発生[%v]", e.Message))
			}
			return ErrorResult(ctx, http.StatusInternalServerError, "INTERNAL SERVER ERROR", fmt.Sprintf("サーバー内でエラーが発生しました[%v]", err))
		}
		return nil
	}
}

// RegisterHandler ハンドラーをセットする
func (handler APIHandler) RegisterHandler(router *echo.Echo) error {
	f, err := middleware.OapiValidatorFromYamlFile("./sample.yml")
	if err != nil {
		return err
	}
	wrapper := oapi.ServerInterfaceWrapper{
		Handler: handler,
	}
	router.Use(errorHandleMiddleware)
	router.Use(f)
	router.GET("/api/healthcheck", wrapper.GetApiHealthcheck)
	router.POST("/api/login", wrapper.PostApiLogin)
	return nil
}
