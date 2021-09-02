package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"sync"
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
	req, _ := http.NewRequest("GET", "http://localhost:8080/healthcheck", nil)
	client := new(http.Client)
	res, err := client.Do(req)
	//res, err := http.Get("http://localhost:8080/healthcheck")
	if err != nil {
		fmt.Println("通信エラー：", err)
		return
	}
	defer res.Body.Close()
	var health HealthCheck
	err = json.NewDecoder(res.Body).Decode(&health)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ヘルスチェックの結果:", health)
}

func execSample1_1() {
	for i := 0; i < 10; i++ {
		sample1()
		time.Sleep(time.Second)
	}
}

func execSample1_2() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			sample1()
			time.Sleep(time.Second)
		}()
	}
	wg.Wait()
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

type RequestAuth struct {
	Token string `json:"token"`
}

func sample3() {
	d, _ := json.Marshal(RequestAuth{
		Token: "TOKEN HOGE",
	})
	reader := bytes.NewReader(d)
	// リクエストを作成
	req, _ := http.NewRequest("POST", "http://localhost:8080/login", reader)
	req.Header = http.Header{
		"Content-Type":  []string{"application/json"},
		"Authorization": []string{"Basic $(echo -n $username:$password | openssl base64)"},
	}

	// リクエスト用クライアント作成
	jar, _ := cookiejar.New(nil)
	// Cookieを保存する
	client := &http.Client{Jar: jar}
	res, err := client.Do(req)
	fmt.Println("LOGIN後:", res.Cookies())
	if err != nil {
		fmt.Println(err)
		return
	}
	req, _ = http.NewRequest("GET", "http://localhost:8080/user", nil)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println("", err)
		return
	}
	time.Sleep(1000)
	req, _ = http.NewRequest("GET", "http://localhost:8080/logout", nil)
	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("LOGOUT後:", res.Cookies())
	req, _ = http.NewRequest("GET", "http://localhost:8080/user", nil)
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
	/*	fmt.Println("---sample1---")
		sample1()*/
	fmt.Println("---sample3---")
	sample3()
	/*	fmt.Println("---sample2---")
		sample2()*/
	/*fmt.Println("---sample3---")
	sample3()
	fmt.Println("---headerSample---")
	headerSample()*/
}
