# Указываем базовый образ
FROM golang:1.23.1 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum файлы
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем остальные файлы приложения
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

RUN ls -la .

# Создаем финальный образ
FROM alpine:latest

WORKDIR /root/

# Копируем бинарный файл из предыдущего образа
COPY --from=builder /app/app .

# Запускаем приложение
CMD ["./app"]
