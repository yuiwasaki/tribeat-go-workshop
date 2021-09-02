package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

func TestPostApiLogin(t *testing.T) {
	str := `{"email":"hoge@hoge.com", "password": "123456"}`
	r := strings.NewReader(str)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", r)
	req.Header = http.Header{"Content-Type": []string{"application/json"}}
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := APIHandler{}
	// ここからテスト
	err := handler.PostApiLogin(c)
	if err != nil {
		t.Error(err)
	}
	if rec.Code != http.StatusCreated {
		t.Error("ステータスコードが正しくない", rec.Code)
	}
	var result oapi.ResultOK
	err = json.NewDecoder(rec.Body).Decode(&result)
	if err != nil {
		t.Error("エラーが発生する", err)
	}
	if !result.Result {
		t.Error("戻り値が正しくない")
	}
}

func BenchmarkXxx(b *testing.B) {
	base := []string{}
	for i := 0; i < b.N; i++ {
		// 都度append
		base = append(base, fmt.Sprintf("no%d", i))
	}
}
