# Task Manager API

Простое API для управления фоновыми задачами. Сервер хранит данные в памяти (при перезапуске все задачи сбрасываются).

## 📌 Особенности
- Легковесное решение для управления задачами
- Минимальные требования к ресурсам
- Простая интеграция
- Данные хранятся только в памяти

## 🚀 Быстрый старт

### Требования
- Go версии 1.20 или выше
- Утилита curl для тестирования (опционально)

### Установка
```bash
git clone https://github.com/Dmitriy4565/task-api/
cd task-api
git checkout master
go mod download
go run main.go
```
## 📚 API Endpoints
### 🆕 Создать задачу
**Метод:** `POST /api/v1/tasks`  
**Тело запроса (JSON):**
```json
{
    "data": "Любой файл для обработки"
}
```
### Пример ответа
```json
{
    "data": {
        "id": "aa49a0d2-4b28-491e-914d-b6b02f3fca0b",
        "status": "pending",
        "created_at": "2025-06-24T19:49:19.338204963+03:00"
    },
    "success": true
}
```
### Проверить статус
**Метод:** `GET /api/v1/tasks/{id}`  
### Пример ответа
```json
{
    "data": {
        "id": "aa49a0d2-4b28-491e-914d-b6b02f3fca0b",
        "status": "completed",
        "created_at": "2025-06-24 19:49",
        "duration": "2 sec",
        "result": "task completed successfully",
        "progress": 97
    },
    "success": true
}
```
### ❌ Удалить задачу
**Метод:** `DELETE /api/v1/tasks/{id}`  
### Ответ: HTTP 204 No Content
