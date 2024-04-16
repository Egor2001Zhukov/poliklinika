package handlers

import (
	"common_go/web/request"
	"common_go/web/response"
	"fmt"
	"io"
	"net/http"
	"poliklinika_gateway/app/settings"
)

func ServerHandler(request *request.Request, response *response.Response) {
	res, err := http.Post(settings.ServerURL, "application/json", request.Body)
	if err != nil {
		panic(fmt.Errorf("ошибка при отправке запроса: %v", err))
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			panic(fmt.Errorf("ошибка при закрытии тела ответа: %v", err))
		}
	}(res.Body)

	// Чтение ответа
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(fmt.Errorf("ошибка при чтении ответа: %v", err))
	}
	response.WriteBody(body)
}

func AuthHandler(req *request.Request, res *response.Response) {

}
