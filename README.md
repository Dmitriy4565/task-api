
# Task Manager API

Простое API для управления фоновыми задачами. Сервер хранит данные в памяти (при перезапуске все задачи сбрасываются).

## API Endpoints

### Создать задачу
**POST /api/v1/tasks**  
Request:
```json
{"data": "любые_данные_для_обработки"}
Проверить статус
GET /api/v1/tasks/{id}
Response:

json
{"id": "123", "status": "processing", "created_at": "2025-06-15 14:30", "duration": "125 sec", "progress": 45.5}
Удалить задачу
DELETE /api/v1/tasks/{id}

Примеры
Создание:

bash
curl -X POST -H "Content-Type: application/json" -d '{"data":"test"}' http://localhost:8080/api/v1/tasks
Проверка:

bash
curl http://localhost:8080/api/v1/tasks/123
Удаление:

bash
curl -X DELETE http://localhost:8080/api/v1/tasks/123
Форматы
Дата: ГГГГ-ММ-ДД ЧЧ:ММ

Время выполнения: 125 sec

Прогресс: 0-100%
