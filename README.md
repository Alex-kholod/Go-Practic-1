# Практическое занятие №14

## Тема: Оптимизация запросов к БД. Использование connection pool

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz14 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
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
1. Добавили в проект репозиторий для работы с базой данных Postgres. Саму базу развернули, используя Docker.
Конфиг пула:
- MaxConns = 20
- MinConns = 5
- MaxConnLifetime = 1h
- MaxConnIdleTime = 10m
3. Создали таблицу Notes в базе данных.
   <img width="974" height="846" alt="image" src="https://github.com/user-attachments/assets/c480f2ce-6f29-44c9-a227-30c91155a0f2" />

4. Создали частичный индекс для таблицы заметок, чтобы искать по заголовку.
   <img width="974" height="720" alt="image" src="https://github.com/user-attachments/assets/2169cca7-503c-453e-830c-2f9979925d45" />

6. Создали индекс для keyset пагинации.
   <img width="974" height="748" alt="image" src="https://github.com/user-attachments/assets/57843061-3253-49a1-a097-cfd43e1d5c33" />

8. Заполнили таблицу заметок тестовыми данными.
   <img width="974" height="822" alt="image" src="https://github.com/user-attachments/assets/b0f520d2-8eac-47d8-a4b2-5b1d49b706a8" />

10. SQL запросы
   - Исходный запрос с OFFSET
     <img width="974" height="652" alt="image" src="https://github.com/user-attachments/assets/39551ee1-2531-4623-b7c7-0fb749c4e8b2" />

   - Запрос с keyset-пагинацией
     <img width="974" height="763" alt="image" src="https://github.com/user-attachments/assets/cf37d49f-6c58-4e38-b999-9ea46f4c68e7" />

#### Вывод: keyset производительнее (время выполнения 0,022 ms против 4,6 ms у OFFSET)

7. Использование batching. Для получения сразу нескольких заметок по ID
   <img width="974" height="655" alt="image" src="https://github.com/user-attachments/assets/b2638e3a-21b1-408e-af77-df44b880b2f1" />
   
### Нагрузочные тесты
1. OFFSET пагинация
```bash
hey -n 2000 -c 50 "http://localhost:8080/api/v1/notes?limit=20&offset=40000"
```
**RPS:** 1359.8
**p95:** 0.0453s
**p99:** 0.0645s
**Error rate:** 0%
<img width="974" height="773" alt="image" src="https://github.com/user-attachments/assets/2fcad0ec-db16-4330-bea5-3b1af44e9975" />

2. Keyset пагинация
```bash
hey -n 2000 -c 50 "http://localhost:8080/api/v1/notes?mode=keyset&limit=20"
```
**RPS:** 6220.2
**p95:** 0.0109s
**p99:** 0.0275s
**Error rate:** 0%
<img width="974" height="769" alt="image" src="https://github.com/user-attachments/assets/ae389273-d3c8-4e13-ac66-802a40e96792" />

1. Batch запрос
```bash
hey -n 5000 -c 50 "http://localhost:8080/api/v1/notes/batch?ids=1,2,3,4,5"
```
**RPS:** 7,309.77
**p95:** 0.0085s
**p99:** 0.0214s
**Error rate:** 0%
<img width="974" height="861" alt="image" src="https://github.com/user-attachments/assets/6283010d-a756-4270-8f00-662f84bd5c62" />

### Выводы
- Keyset-пагинация дала наибольший прирост производительности по сравнению с OFFSET.
- Batching позволяет получить набор сущностей за один запрос вместо множества запросов получения по 1 элементу.
- Connection pool позволяет контролировать число соединений с базой и стабилизировать работу под нагрузкой сервера бд.

   
### Структура проекта
<img width="749" height="1217" alt="image" src="https://github.com/user-attachments/assets/6f8bbd79-5a07-41ea-82db-1301d2e7542c" />
