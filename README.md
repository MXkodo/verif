# Authentication Service

## Описание

Этот проект представляет собой сервис аутентификации, реализованный на языке Go. Он использует JWT для Access токенов и bcrypt для хранения Refresh токенов. Сервис хранит данные в базе данных PostgreSQL и работает в контейнерах Docker.

## Технологии

- **Go**: Язык программирования
- **JWT**: Для Access токенов
- **PostgreSQL**: Для хранения данных
- **Docker**: Для контейнеризации приложения

## Установка и запуск

### Сборка Docker-образа

Выполните команду для сборки Docker-образа:

```bash
docker build -t auth-service .
```

Выполните команду для запуска Docker-образа:

```bash
docker-compose up
```

## Доступные API

### Генерация токенов

#### URL: POST http://localhost:8080/tokens/generate

```json
{
  "userID": "123e4567-e89b-12d3-a456-426614174000",
  "ip": "127.0.0.1"
}
```

### Обновление токенов

#### URL: POST http://localhost:8080/tokens/refresh

```json
{
  "refreshToken": "e2tLTPwGAGONZyTc9C4omUE4udiPxkLXDf33psDRcFY=",
  "ip": "127.0.0.1"
}
```
