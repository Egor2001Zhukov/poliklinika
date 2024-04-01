package main

import (
	"fmt"
	"net/http"
)

func getResponse(w http.ResponseWriter, r *http.Request) {
	// Получаем путь запроса
	path := r.URL.Path

	// Ваша логика для получения ответа на запрос
	response := fmt.Sprintf("Response for path: %s", path)

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func main() {

	common_middlewares =

	// Создаем маршрутизатор
	mux := http.NewServeMux()

	// Обработка всех запросов с помощью одного обработчика
	mux.HandleFunc("/", getResponse)

	// Запускаем сервер
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
