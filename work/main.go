package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/models"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

type MemberId struct {
	MemberId string `json:"member_id"`
}

type APIHandler struct {
	db *models.Model
}

func main() {
	/*db, err := models.NewModel()
	if err != nil {
		fmt.Println(err)
		return
	}*/
	e := echo.New()
	e.Use(errorMeesageMiddleware)
	f, err := middleware.OapiValidatorFromYamlFile("./sample.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	e.Use(f)
	handler := APIHandler{
		//db: db,
	}
	RegisterHandlers(e, handler)
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", 3003)))
}

// ユーザー一覧取得
// (GET /api/client/users)
func (api APIHandler) GetApiClientUsers(ctx echo.Context, params oapi.GetApiClientUsersParams) error {
	/*if err == NotFound {
		// 404
	}*/
	result := oapi.ResultUsers{}
	return ctx.JSON(http.StatusOK, result)
}

// ユーザー新規登録
// (POST /api/client/users)
func (api APIHandler) PostApiClientUsers(ctx echo.Context) error {
	var req oapi.RequestUser
	err := ctx.Bind(&req)
	if err != nil {
		fmt.Println(err)
		return ErrorResult(ctx, http.StatusBadRequest, "正しくない", "リクエストが正しくない")
	}
	d, _ := json.Marshal(req)
	fmt.Println(string(d))
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
	model := models.NewGroupModel(api.db)
	group := oapi.Group{
		Id:   id.String(),
		Name: req.Name,
	}
	err = model.Save(models.Group{
		Group: group,
	})
	if err != nil {
		fmt.Println(err)
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
		req := ctx.Request()
		cid, ok1 := req.Header["Cid"]
		apptoken, ok2 := req.Header["Apptoken"]
		if !ok1 || !ok2 {
			// エラー
			return ErrorResult(ctx, http.StatusUnauthorized, "トークンエラー", "トークンがありません")
		}
		memberId := ""
		if req.Method == "POST" || req.Method == "PUT" || req.Method == "PATCH" {
			body, err := ioutil.ReadAll(ctx.Request().Body)
			if err != nil {
				return ErrorResult(ctx, http.StatusUnauthorized, "トークンエラー", "メンバーが存在しません")
			}
			var mem MemberId
			json.Unmarshal(body, &mem)
			memberId = mem.MemberId
			ctx.Request().Body = ioutil.NopCloser(bytes.NewBuffer(body))
		} else {
			vals := ctx.QueryParams()
			if mems, ok := vals["member_id"]; ok {
				memberId = mems[0]
			} else {
				// メンバーIDが存在しない
				return ErrorResult(ctx, http.StatusUnauthorized, "トークンエラー", "メンバーが存在しません")
			}
		}
		fmt.Println("CID:", cid[0])
		fmt.Println("APPTOKEN:", apptoken[0])
		fmt.Println("MemberId:", memberId)
		// ログインチェック
		fmt.Println("トークンが正しい")
		return next(ctx)
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
	model := models.NewGroupModel(api.db)
	groups, err := model.All()
	if err != nil {
		fmt.Println(err)
		return ErrorResult(ctx, http.StatusInternalServerError, "サーバーエラー", "エラー発生")
	}
	gs := make([]oapi.Group, len(groups))
	for i, g := range groups {
		gs[i] = g.Group
	}
	result := oapi.ResultGroupsGet{
		Count: len(groups),
		Data:  gs,
	}
	return ctx.JSON(http.StatusOK, result)
}

func RegisterHandlers(router *echo.Echo, handler APIHandler) {
	wrapper := oapi.ServerInterfaceWrapper{
		Handler: handler,
	}
	router.GET("/api/healthcheck", wrapper.GetApiHealthcheck)
	router.POST("/api/login", wrapper.PostApiLogin)
	clientGroup := router.Group("/api/client", authCheckMiddleware)
	clientGroup.GET("/users", wrapper.GetApiClientUsers)
	clientGroup.POST("/users", wrapper.PostApiClientUsers)
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
