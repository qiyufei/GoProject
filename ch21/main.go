package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8080", nil)

	// 下载&使用Gin框架
}

type user struct {
	Id   int
	Name string
}

var users = []user{
	{1, "yufei"},
	{2, "na"},
	{3, "nana"},
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		re, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(re)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "{\"message\": \"not found\"}")
	}
}
