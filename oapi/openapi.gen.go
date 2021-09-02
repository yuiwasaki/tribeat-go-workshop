// Package oapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.8.2 DO NOT EDIT.
package oapi

import (
	"fmt"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// グループ一覧取得
	// (GET /api/client/groups)
	GetApiClientGroups(ctx echo.Context) error
	// グループ作成
	// (POST /api/client/groups)
	PostApiClientGroups(ctx echo.Context) error
	// ユーザー一覧取得
	// (GET /api/client/users)
	GetApiClientUsers(ctx echo.Context, params GetApiClientUsersParams) error
	// ユーザー新規登録
	// (POST /api/client/users)
	PostApiClientUsers(ctx echo.Context) error
	// ユーザー一覧取得
	// (GET /api/client/users/{userId})
	GetApiClientUsersUserId(ctx echo.Context, userId UserId) error
	// ヘルスチェック
	// (GET /api/healthcheck)
	GetApiHealthcheck(ctx echo.Context) error
	// ログイン
	// (POST /api/login)
	PostApiLogin(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetApiClientGroups converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiClientGroups(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiClientGroups(ctx)
	return err
}

// PostApiClientGroups converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiClientGroups(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostApiClientGroups(ctx)
	return err
}

// GetApiClientUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiClientUsers(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetApiClientUsersParams
	// ------------- Required query parameter "member_id" -------------

	err = runtime.BindQueryParameter("form", true, true, "member_id", ctx.QueryParams(), &params.MemberId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter member_id: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiClientUsers(ctx, params)
	return err
}

// PostApiClientUsers converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiClientUsers(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostApiClientUsers(ctx)
	return err
}

// GetApiClientUsersUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiClientUsersUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId UserId

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiClientUsersUserId(ctx, userId)
	return err
}

// GetApiHealthcheck converts echo context to params.
func (w *ServerInterfaceWrapper) GetApiHealthcheck(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApiHealthcheck(ctx)
	return err
}

// PostApiLogin converts echo context to params.
func (w *ServerInterfaceWrapper) PostApiLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostApiLogin(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/client/groups", wrapper.GetApiClientGroups)
	router.POST(baseURL+"/api/client/groups", wrapper.PostApiClientGroups)
	router.GET(baseURL+"/api/client/users", wrapper.GetApiClientUsers)
	router.POST(baseURL+"/api/client/users", wrapper.PostApiClientUsers)
	router.GET(baseURL+"/api/client/users/:userId", wrapper.GetApiClientUsersUserId)
	router.GET(baseURL+"/api/healthcheck", wrapper.GetApiHealthcheck)
	router.POST(baseURL+"/api/login", wrapper.PostApiLogin)

}

