package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/yuiwasaki/tribeat-go-workshop/oapi"
)

func TestPostApiLogin(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	handler := APIHandler{}
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
