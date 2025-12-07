# Практическое занятие №11

## Тема: Проектирование REST API (CRUD для заметок). Разработка структуры

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz11 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./cmd/api
```

### Ход работы
1. Реализовали API для выполнения CRUD-операций с заметками. 
2. Сделали слоистую архитектуру в проекте, выделив слой обработчиков, сервисов и репозитория.
3. Запустили код и выполнили запросы для проверки API.
4. Выполнение CRUD-операций:
   - POST /api/v1/notes (Создание заметки)
     <img width="974" height="636" alt="image" src="https://github.com/user-attachments/assets/957b6a28-5382-4136-8285-eaa97320f0d9" />

   - GET /api/v1/notes (Получение списка заметок)
     <img width="974" height="652" alt="image" src="https://github.com/user-attachments/assets/3045cdf5-c31d-4db0-bbef-a9585d8f80cd" />

   - GET /api/v1/notes/{id} (Получение заметки по ID)
     <img width="974" height="674" alt="image" src="https://github.com/user-attachments/assets/c0597070-3c61-4459-a91b-de1f9a594bb8" />

   - PUT /api/v1/notes/{id} (Обновление заметки по ID)
     <img width="974" height="646" alt="image" src="https://github.com/user-attachments/assets/22745ece-0cbf-4d67-b4b9-ac2c9a677ebe" />

   - DELETE /api/v1/notes/{id} (Удаление заметки по ID)
     <img width="974" height="680" alt="image" src="https://github.com/user-attachments/assets/2fe4d857-2483-48cf-b7a5-d299679d9eb2" />
     <img width="974" height="646" alt="image" src="https://github.com/user-attachments/assets/4bb10071-1c33-4a40-a165-0a0fb6671c23" />

   
### Структура проекта
<img width="553" height="845" alt="image" src="https://github.com/user-attachments/assets/12231bbc-a6bc-404d-aec1-da63eb6e2ee0" />


