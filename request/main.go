package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// HealthCheck これはヘルスチェックの戻り
type HealthCheck struct {
	// Result 成功:true 失敗:false
	Result bool `json:"result"`
}

// httpレスポンスのBodyをそのままJSONに変換する
func sample1() {
	// リクエストを作成
	req, _ := http.NewRequest("GET", "http://localhost:4010/api/healthcheck", nil)
	req.Header.Set("Content-Type", "application/json")
	// リクエスト用クライアント作成
	client := new(http.Client)
	fmt.Println(req.Header)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	var health HealthCheck
	err = json.NewDecoder(res.Body).Decode(&health)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(health.Result)
}

func sample2() {
	res, err := http.Get("http://localhost:4010/api/healthcheck")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	var health HealthCheck
	err = json.NewDecoder(res.Body).Decode(&health)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(health)
}

func sample3() {
	// リクエストを作成
	req, _ := http.NewRequest("POST", "http://localhost:4010/api/login", nil)
	// リクエスト用クライアント作成
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, _ = http.NewRequest("GET", "http://localhost:4010/api/users", nil)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	time.Sleep(1000)
	req, _ = http.NewRequest("GET", "http://localhost:4010/api/logout", nil)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Cookies())
	req, _ = http.NewRequest("GET", "http://localhost:4010/api/users", nil)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Cookies())
}

// ヘッダーサンプル
func headerSample() {
	req, _ := http.NewRequest("GET", "http://localhost:4010/api/healthcheck", nil)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	fmt.Println("ステータスコード：", res.StatusCode)
	for key, val := range res.Header {
		fmt.Println("KEY:", key, "VAL:", val)
	}
	fmt.Println("Content-Type:", res.Header.Get("content-type"))
}

func main() {
	fmt.Println("---sample1---")
	sample1()
	fmt.Println("---sample2---")
	sample2()
	/*fmt.Println("---sample3---")
	sample3()
	fmt.Println("---headerSample---")
	headerSample()*/
}
