# Используем официальный образ Go
FROM golang:latest as build

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Скачиваем зависимости
RUN go mod tidy

# Собираем приложение
RUN go build -o myapp .

# Финальный контейнер
FROM golang:latest

WORKDIR /app

# Копируем скомпилированное приложение
COPY --from=build /app/myapp .

# Открываем порт
EXPOSE 8070

# Запускаем приложение
CMD ["/app/myapp"]