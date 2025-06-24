# Task Manager API

Простое API для управления фоновыми задачами. Сервер хранит данные в памяти (при перезапуске все задачи сбрасываются).

## Как запустить

1. Установите Go (версия 1.20+)
2. Склонируйте репозиторий
3. Выполните в терминале:

```bash
go mod download
go run main.go
Сервер запустится на http://localhost:8080

API Endpoints
Создать задачу
text
POST /api/v1/tasks
Тело запроса (JSON):

json
{
    "data": "любые_данные_для_обработки"
}
Проверить статус
text
GET /api/v1/tasks/{id}
Пример ответа:

json
{
    "id": "123",
    "status": "processing",
    "created_at": "2025-06-15 14:30",
    "duration": "125 sec",
    "progress": 45.5
}
Удалить задачу
text
DELETE /api/v1/tasks/{id}
Форматы данных
Дата: ГГГГ-ММ-ДД ЧЧ:ММ (например 2025-06-15 14:30)

Продолжительность: в секундах (125 sec)

Прогресс: от 0 до 100%

Примеры использования
Создаем задачу:

bash
curl -X POST -H "Content-Type: application/json" -d '{"data":"test"}' http://localhost:8080/api/v1/tasks
Проверяем статус:

bash
curl http://localhost:8080/api/v1/tasks/123
