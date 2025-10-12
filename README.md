# Практическое занятие №5

## Тема: Подключение к PostgreSQL через database/sql. Выполнение простых запросов (INSERT, SELECT)

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz5 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./app
```

### Ход работы
1. Запустили Postgres в docker контейнере. Создали базу данных с названием "todo". Подключились к базе и создали таблицу задач.
   <img width="974" height="567" alt="image" src="https://github.com/user-attachments/assets/3f46460d-3e81-47b5-ad0e-edee87d28163" />

3. Подключились в коде к базе и добавили задачи.
   <img width="974" height="353" alt="image" src="https://github.com/user-attachments/assets/3c0c3661-ccc3-40c7-a36c-f9e93da966dc" />
   <img width="974" height="357" alt="image" src="https://github.com/user-attachments/assets/a30c4b2b-e7b0-4f79-9e1d-3d5154ad9010" />

4. Добавлена функция вывода задач по фильтру «выполнено».
   <img width="974" height="404" alt="image" src="https://github.com/user-attachments/assets/494052fe-750b-4f83-838d-dbf2cdcac24a" />

5. Добавлена функция вывода задач по ID (На примере задача с ID=1).
   <img width="974" height="113" alt="image" src="https://github.com/user-attachments/assets/346f0433-9e06-4103-9bea-dbea5d7e1b76" />

6. Добавлена функция добавления списка задач через транзакцию.
   <img width="450" height="111" alt="image" src="https://github.com/user-attachments/assets/d024a9cf-09e2-4ce4-af7e-4292d49aed26" />


## Структура проекта
<img width="366" height="401" alt="image" src="https://github.com/user-attachments/assets/2e961ae8-c588-4bcd-ad46-e4e6c501abd8" />
