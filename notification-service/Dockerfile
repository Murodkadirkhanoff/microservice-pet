# Этап сборки
FROM golang:latest as build

WORKDIR /app

# Копируем файлы проекта
COPY . .

# Скачиваем зависимости
RUN go mod download

# Собираем приложение
RUN go build -o /myapp .

# Финальный контейнер
FROM debian:latest

WORKDIR /app

# Копируем бинарник из этапа сборки
COPY --from=build /myapp /app/myapp

# Проверяем, что файл существует
RUN ls -lah /app/myapp

# Даем права на выполнение
RUN chmod +x /app/myapp

# Открываем порт
EXPOSE 8081

# Запускаем приложение
CMD ["/app/myapp"]
