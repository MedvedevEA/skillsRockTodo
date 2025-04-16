# skillsRockTodo
## Skills Rock Todo List
Базовый HTTP-сервер с маршрутизацией. Хранить данные в памяти реализовано на основании in-memory хранилища. 
## Запуск проекта
1. Клонируйте репозиторий
2. Создайте файл /config/local.yml по образцу
3. Выполните миграцию БД 
```
go run cmd/migrator/main.go -source-path migration/ -database-url "postgres://postgres:1234@localhost:5433/temp?sslmode=disable"
```
4. Запустите проект
```
go run cmd/todo/main.go -config config/local.yml
```
