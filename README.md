# Практическое занятие №1

## Тема: Установка и настройка окружения Go

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz1 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./cmd/server
```
Альтернативный запуск с помощью файла helloapi.exe

### Принцип работы
После запуска приложения можно использовать его endpoint's для взаимодействия.
По умолчанию запуск на порту 8080. Возможен запуск на другом порту, через настройку переменной окружения.
<img width="906" height="170" alt="image" src="https://github.com/user-attachments/assets/2fec9b5e-96ca-4a14-ab20-bb3eb884f239" />

#### Поддерживаемые endpoint's:
## 1. GET /hello
   #### Формат ответа:
   <img width="974" height="345" alt="image" src="https://github.com/user-attachments/assets/82a55bc2-fc69-43e5-ae59-fd7e6c42756a" />
   
## 2. GET /user
   #### Формат ответа:
   <img width="974" height="349" alt="image" src="https://github.com/user-attachments/assets/927b8915-d0c9-4b7d-a91e-68bcc7c8e66b" />

## 4. GET /health
   #### Формат ответа:
   <img width="974" height="343" alt="image" src="https://github.com/user-attachments/assets/33e30903-c4a5-40a7-bd43-67dab5b44981" />

# Структура проекта:

## Описание структуры:
- cmd/server - Основное приложение
- go.mod - Файл модуля Go с зависимостями
- go.sum - Файл версий зависимостей проекта
