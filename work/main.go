package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/models"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

type APIHandler struct {
	db *models.Model
}

func main() {
	fmt.Println("START")
	fmt.Println("CONNECT DB")
	db, err := models.NewModel()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("DB CONNECTED")

	e := echo.New()
	e.Use(errorMeesageMiddleware)
	f, err := middleware.OapiValidatorFromYamlFile("./sample.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	e.Use(f)
	handler := APIHandler{
		db: db,
	}
	RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 3003)))
}

// ユーザー一覧取得
// (GET /api/client/users)
func (api APIHandler) GetApiClientUsers(ctx echo.Context) error {
	/*if err == NotFound {
		// 404
	}*/
	result := oapi.ResultUsers{}
	return ctx.JSON(http.StatusOK, result)
}

// ユーザー新規登録
// (POST /api/client/users)
func (api APIHandler) PostApiClientUsers(ctx echo.Context) error {
	return nil
}

// ユーザー一覧取得
// (GET /api/client/users/{userId})
func (api APIHandler) GetApiClientUsersUserId(ctx echo.Context, userId oapi.UserId) error {
	user := ctx.Get("user").(oapi.User)
	fmt.Println("USER INFO:", user)
	return ctx.JSON(http.StatusOK, user)
}

// グループ作成
// (POST /api/client/groups)
func (api APIHandler) PostApiClientGroups(ctx echo.Context) error {
	var req oapi.RequestGroupsPost
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return ErrorResult(ctx, http.StatusBadRequest, "正しくない", "リクエストが正しくない")
	}

	id, _ := uuid.NewRandom()
	group := oapi.Group{
		Id:   id.String(),
		Name: req.Name,
	}
	now := time.Now()
	tx := api.db.Create(&models.Group{
		Group:     group,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return ErrorResult(ctx, http.StatusInternalServerError, "保存失敗", "サーバー内エラーが発生しました")
	}
	return ctx.JSON(http.StatusCreated, oapi.ResultGroupsPost{
		Result: true,
		Data:   group,
	})
}

func checkUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := ctx.Param("userId")
		if userId != "aaa" {
			return ErrorResult(ctx, http.StatusNotFound, "存在しない", "ユーザーが存在しません")
		}
		ctx.Set("user", oapi.User{
			Id:   "aaa",
			Name: "IWASAKI",
		})
		return next(ctx)
	}
}

var (
	ERR_NO_AUTH_TOKEN = fmt.Errorf("NO AUTH")
	ERR_INVALID_TOKEN = fmt.Errorf("INVALID TOKEN")
)

func authCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// ログインチェック
		if token, ok := ctx.Request().Header["Authorization"]; !ok {
			return ERR_NO_AUTH_TOKEN
			//return ErrorResult(ctx, http.StatusUnauthorized, "トークンエラー", "トークンがありません")
		} else if token[0] != "token" {
			return ERR_INVALID_TOKEN
			//return ErrorResult(ctx, http.StatusUnauthorized, "トークンエラー", "トークンが正しくありません")
		}
		fmt.Println("トークンが正しい")
		err := next(ctx)
		return err
	}
}

/*
func maagerRoleCheckMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 管理者権限確認処理
		if err != nil {
			return ErrorResult(ctx, http.StatusInternalServerError, "", "")
		}
		return next(ctx)
	}
}*/

func errorMeesageMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// 処理
		err := next(ctx)
		if err != nil {
			if err == ERR_INVALID_TOKEN {
				return ErrorResult(ctx, http.StatusBadRequest, "AAA", "トークンが正しくない")
			} else if err == ERR_NO_AUTH_TOKEN {
				return ErrorResult(ctx, http.StatusBadRequest, "AAA", "トークンがない")
			}
			return ErrorResult(ctx, http.StatusBadRequest, "AAA", fmt.Sprintf("エラーが発生[%v]", err))
		}
		return nil
	}
}

/*
func dumpMessage(c echo.Context, reqBody, resBody []byte) {
	//fmt.Println(c.Request().Header)
	//fmt.Println(c.Request().Header.Get("User-Agent"))

	//contentType := c.Request().Header.Get("Content-Type")
	//fmt.Printf("Request Content-Type: %s\n", contentType)
	//if strings.Contains(contentType, "application/json") {
	//	fmt.Printf("Request Body: %v\n", string(reqBody))
	//}
	//contentType = c.Response().Header().Get("Content-Type")
	//fmt.Printf("Response Content-Type: %s\n", contentType)
	//if c.Response().Header().Get("Content-Type") == "application/json; charset=UTF-8" {
	//	fmt.Printf("Response Body: %v\n", string(resBody))
	//}
	return func(ctx echo.Context) error {
		fmt.Println(ctx.Request().Header)
		err := next(ctx)
	}
}*/

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

// GetApiHealthcheck converts echo context to params.
func (api APIHandler) GetApiHealthcheck(ctx echo.Context) error {
	//if err != nil {
	//	return ErrorResult(ctx, http.StatusInternalServerError, "ERROR123", "HOGE")
	//}
	result := oapi.ResultOK{
		Result: true,
	}
	return ctx.JSON(http.StatusOK, result)
}

// ログイン
// (POST /api/login)
func (api APIHandler) PostApiLogin(ctx echo.Context) error {
	var v oapi.RequestLogin
	err := ctx.Bind(&v)
	if err != nil {
		return ErrorResult(ctx, http.StatusUnauthorized, "ERROR", err.Error())
	}
	fmt.Println(v)
	result := oapi.ResultOK{
		Result: true,
	}
	return ctx.JSON(http.StatusOK, result)
}

// グループ一覧取得
// (GET /api/client/groups)
func (api APIHandler) GetApiClientGroups(ctx echo.Context) error {
	_, err := api.db.SelectGroups()
	if err != nil {
		fmt.Println(err)
		return ErrorResult(ctx, http.StatusInternalServerError, "サーバーエラー", "エラー発生")
	}
	return nil
}

func RegisterHandlers(router *echo.Echo, handler APIHandler) {
	wrapper := oapi.ServerInterfaceWrapper{
		Handler: handler,
	}
	router.GET("/api/healthcheck", wrapper.GetApiHealthcheck)
	router.POST("/api/login", wrapper.PostApiLogin)
	clientGroup := router.Group("/api/client", authCheckMiddleware)
	clientGroup.GET("/users", wrapper.GetApiClientUsers)
	userGroup := clientGroup.Group("/users/:userId", checkUserMiddleware)
	userGroup.GET("", wrapper.GetApiClientUsersUserId)
	clientGroup.GET("/groups", wrapper.GetApiClientGroups)
	clientGroup.POST("/groups", wrapper.PostApiClientGroups)
	//userGroup.DELETE("", wrapper.DeleteApiClientUsersUserId)
	//userGroup.PUT("", wrapper.PutApiClientUsersUserId)
	/*
		clientGroup := router.Group("/api/client", authCheckMiddleware)
		// /api/client/user
		clientGroup.GET("/user", getUser)

		managerGroup := router.Group("/api/manager", authCheckMiddleware)
		managerGroup = managerGroup.Group("", managerCheckMiddleware)
		managerGroup.GET("/users", getUsers)
		managerUserGroup := managerGroup.Group("/users/:userId", findUserMiddleware)
		// /api/manager/users/:userId
		managerUserGroup.GET("", getUserInfo)
		// /api/manager/users/:userId
		managerUserGroup.PUT("", putUserInfo)
		// /api/manager/users/:userId/status
		managerUserGroup.GET("/status", getUserStatus)*/

}