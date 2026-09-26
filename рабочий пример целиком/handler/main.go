package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Task struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks := []Task{
		{ID: 1, Title: "Сверстать карточку", Done: true},
		{ID: 2, Title: "Написать хендлер", Done: false},
		{ID: 3, Title: "Проверить API", Done: false},
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	req := httptest.NewRequest("GET", "/tasks", nil)
	rec := httptest.NewRecorder()
	tasksHandler(rec, req)
	res := rec.Result()
	fmt.Println("статус:", res.StatusCode)
	fmt.Println("Content-Type:", res.Header.Get("Content-Type"))
	fmt.Println("тело:", rec.Body.String())
}
