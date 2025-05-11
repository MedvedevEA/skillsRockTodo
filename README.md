# skillsRockTodo
## Skills Rock Todo List
Базовый HTTP-сервер с маршрутизацией.
## Запуск проекта
1. Клонируйте репозиторий
2. Создайте файл /config/local.yml по образцу
3. Выполните миграцию БД 
```
go run cmd/migrator/main.go -source-path migration/ -database-url "postgres://postgres:postgres@localhost:5432/db?sslmode=disable"
```
4. Запустите проект
```
go run cmd/todo/main.go -config config/local.yml
```
