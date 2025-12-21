# Практическое занятие №16

## Тема: Интеграционное тестирование API.

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz16 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
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
2. Добавили в архитектуру папку integration, куда поместили файл с интеграционными тестами.
3. Добавили миграцию на создание таблицы Notes в бд при запуске тестов.
   <img width="974" height="578" alt="image" src="https://github.com/user-attachments/assets/6e74e463-5dfd-4579-92ef-b1877751f2dd" />

5. Запустили тесты командой
```bash
go test ./internal/integration -v
```
#### Результаты вполнения тестов - успешно
<img width="822" height="484" alt="image" src="https://github.com/user-attachments/assets/19d1b89b-a9e6-403a-b2f1-1d7c68503991" />

#### Таблица в бд после выполнения
<img width="974" height="586" alt="image" src="https://github.com/user-attachments/assets/9a1f821c-1070-47e5-931a-1054bd0683b7" />


### Структура проекта
<img width="709" height="1288" alt="image" src="https://github.com/user-attachments/assets/47e13b5a-d10b-4960-9278-3a05445b513c" />
