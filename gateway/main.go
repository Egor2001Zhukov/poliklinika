package main

import (
	"auth/app/settings"
	"auth/app/urls"
	"fmt"
	"net/http"
)

func processRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
	f, ok := urls.Urls[r.RequestURI]
	if !ok {
		http.Error(w, "NotFound", http.StatusNotFound)
		return
	}
	handler := http.HandlerFunc(f)
	// Применяем middleware к обработчику
	for _, middleware := range settings.CommonMiddlewares {
		handler = middleware(handler)
	}

	// Обработка запроса
	handler.ServeHTTP(w, r)
}

func main() {

	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Обработка всех запросов с помощью одного обработчика
	mux.HandleFunc("/", processRequest)

	// Запускаем сервер
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
