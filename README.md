# Практическое занятие №8

## Тема: Работа с MongoDB: подключение, создание коллекции, CRUD-операции

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz8 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
4. Запуск MongoDB
```bash
docker compose up -d
```
4. Запуск приложения
```bash
go run ./cmd/api
```

### Принцип работы
После запуска приложения можно использовать его endpoint's для взаимодействия.
#### Поддерживаемые CRUD-операции взаимодействия с Notes, а также другие endpoints:
## 1. GET /
   #### Формат ответа:
   <img width="974" height="486" alt="image" src="https://github.com/user-attachments/assets/6aae1163-43b5-47be-a905-ce4cb8ee7a86" />
   
## 2. POST /
   #### Формат ответа:
   <img width="974" height="476" alt="image" src="https://github.com/user-attachments/assets/26dfb12a-ec0a-4d9a-9c0c-f9f41d0059be" />

## 3. GET /{id}
   #### Формат ответа:
   <img width="974" height="465" alt="image" src="https://github.com/user-attachments/assets/d3914fd2-f0c5-413d-b857-ee5e81f00f54" />

## 4. PATCH /{id}
   #### Формат ответа:
   <img width="974" height="459" alt="image" src="https://github.com/user-attachments/assets/4f185e5e-7349-4962-ab04-2962c8bb1438" />

## 5. DELETE /{id}
   #### Формат ответа:
   <img width="974" height="369" alt="image" src="https://github.com/user-attachments/assets/b9f01fc0-68e5-4e14-937d-5c9c65fcab14" />
   #### После удаления единственной записи:
   <img width="974" height="414" alt="image" src="https://github.com/user-attachments/assets/4cf97be7-7425-4fa8-b4f6-a6f9b8e1ee61" />

## 6. GET /stats (Статистика: количество заметок, средняя/максимальная/минимальная длина content)
   #### Формат ответа:
   <img width="974" height="442" alt="image" src="https://github.com/user-attachments/assets/5b3733cd-309f-4ef7-aeaf-a4c7f4c7b91e" />

## 7. GET /?after={id}&limit={count} (Пагинация)
   #### Формат ответа:
   <img width="974" height="497" alt="image" src="https://github.com/user-attachments/assets/6810d38c-3943-457b-a06b-ba3041b31b8a" />

## 8. POST / (Создание временной заметки с датой удаления)
   #### Формат ответа:
   <img width="974" height="478" alt="image" src="https://github.com/user-attachments/assets/28d7cb87-2b80-4e2f-a447-dab391c3424f" />


   
