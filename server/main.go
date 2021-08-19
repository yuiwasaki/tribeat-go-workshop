package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/yuiwasaki/tribeat-go-workshop/server/types"
)

func handlerHealthCheck(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		header := w.Header()
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Hello", "World")
		header.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		health := types.Result{
			Result: true,
		}
		json.NewEncoder(w).Encode(health)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	health := types.Result{
		Result: false,
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(health)
}

func handlerLogin(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "Auth",
		Value:   "auth",
		Path:    "/",
		Expires: time.Now().Add(time.Hour),
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	result := types.Result{
		Result: true,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func handlerLogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    "Auth",
		Value:   "expired",
		Path:    "/",
		Expires: time.Now(),
	}
	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	result := types.Result{
		Result: true,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("Auth")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c.Value)
	}
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		user := types.ResultUser{
			Result: types.Result{
				Result: true,
			},
			Data: types.User{
				Id:   "HOGE",
				Name: "PIYO",
			},
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
		return
	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		user := types.Result{
			Result: true,
		}
		json.NewEncoder(w).Encode(user)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	health := types.Result{
		Result: false,
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(health)
}

func main() {
	http.HandleFunc("/healthcheck", handlerHealthCheck)
	http.HandleFunc("/user", handlerUser)
	http.HandleFunc("/login", handlerLogin)
	http.HandleFunc("/logout", handlerLogout)
	http.ListenAndServe(":8080", nil)
}
