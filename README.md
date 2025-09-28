# Практическое занятие №3

## Тема: Реализация простого HTTP-сервера на стандартной библиотеке net/http

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz3 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./cmd/server
```

### Принцип работы
После запуска приложения можно использовать его endpoint's для взаимодействия.
#### Поддерживаемые endpoint's:
## 1. GET /health
   #### Формат ответа:
   <img width="974" height="336" alt="image" src="https://github.com/user-attachments/assets/bcfdddda-1233-45e4-91db-5a644b9568ac" />

## 2. POST /tasks
   #### Формат ответа:
   <img width="974" height="519" alt="image" src="https://github.com/user-attachments/assets/435277a7-948f-4c5a-b752-ec35de06f15c" />


## 4. GET /tasks
   #### Формат ответа:
   <img width="974" height="577" alt="image" src="https://github.com/user-attachments/assets/8bb519fd-64e3-4b32-880c-47d13d372669" />


## 5. PATCH /tasks/{id}
   #### Формат ответа:
   <img width="974" height="476" alt="image" src="https://github.com/user-attachments/assets/0f68441d-b313-4988-9c0b-3d5d6232d59d" />


## 6. DELETE /tasks/{id}
   #### Формат ответа:
   <img width="1389" height="586" alt="image" src="https://github.com/user-attachments/assets/79e16dd5-cbc7-4856-932a-c2318fc14c9d" />


### Настройка CORS в проекте:

Результат запроса с fontend:
1. Не успешный
   <img width="974" height="221" alt="image" src="https://github.com/user-attachments/assets/2f7e3b6a-811b-480a-aafc-359b62c0368d" />

2. Успешный
   <img width="927" height="309" alt="image" src="https://github.com/user-attachments/assets/a862bbf2-0310-46c5-adbb-3af31abbdf2a" />

# Запуск UNIT-tests
```bash
go test -v ./internal/api/...
```
Результат вывода:
<img width="974" height="465" alt="image" src="https://github.com/user-attachments/assets/b6b45bff-2d31-42a5-9483-351ff71dbdd1" />

