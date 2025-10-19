# Практическое занятие №6

## Тема: Использование ORM (GORM). Модели, миграции и связи между таблицами

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz6 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
cd .\app\
go run ./cmd/server
```

### Ход работы
1. Запустили Postgres в docker контейнере. Создали базу данных с названием "pz6_gorm". Настроили подключение к базе в коде, используя переменные окружения.
   
2. Запустили код и выполнили запросы
   - /health
     <img width="934" height="195" alt="image" src="https://github.com/user-attachments/assets/f66aa500-1aff-4637-9e3d-44cc99c49e91" />

   - /users /notes /notes/1
     <img width="974" height="374" alt="image" src="https://github.com/user-attachments/assets/260f5d51-cf9d-4117-8c7a-a73736cf9810" />


3. Автоматически с помощью миграций создались таблицы в базе данных и заполнились значениями из запросов.
   <img width="555" height="381" alt="image" src="https://github.com/user-attachments/assets/6800fca3-669d-4c15-8d83-5b123486de3f" />
#### Таблица notes:
   <img width="974" height="244" alt="image" src="https://github.com/user-attachments/assets/2bcbccf0-0422-49da-aa06-7e8dcea6b5fa" />
#### Таблица tags:
   <img width="974" height="339" alt="image" src="https://github.com/user-attachments/assets/50811ac2-b42d-4089-9db1-c937d0592277" />
#### Таблица users:
   <img width="974" height="258" alt="image" src="https://github.com/user-attachments/assets/cb0d21a8-5818-4949-ab94-41314120dfb5" />
#### Таблица note_tags:
   <img width="633" height="400" alt="image" src="https://github.com/user-attachments/assets/40099d91-ed43-4072-8993-b40da7126227" />

## Структура проекта
<img width="513" height="419" alt="image" src="https://github.com/user-attachments/assets/1bdd6644-757f-4792-9860-3a184ffced2d" />
