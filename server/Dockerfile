# Используем официальный образ Golang
FROM golang:latest

# Устанавливаем рабочую директорию внутри контейнера

WORKDIR /poliklinika

COPY ./common_go common_go/
# Копируем содержимое текущей директории внутрь контейнера
COPY ./server server/

# Компилируем Go приложение
RUN cd server && go build -o server .

# Запускаем приложение при старте контейнера
CMD ["./server/server"]
