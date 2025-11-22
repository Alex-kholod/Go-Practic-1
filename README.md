# Практическое занятие №9

## Тема: Реализация регистрации и входа пользователей. Хэширование паролей с bcrypt

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz9 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
cd .\app\
go run ./cmd/api
```

### Ход работы
1. Запустили Postgres в docker контейнере. Создали базу данных с названием "pz9". Настроили подключение к базе в коде, используя переменные окружения.
   
2. Запустили код и выполнили запросы на регистрацию и авторизацию
   - /auth/register
      Регистрация
      <img width="974" height="466" alt="image" src="https://github.com/user-attachments/assets/550e1d9c-867e-4c92-aa9b-6c8a7d2dd839" />
      Повторная регистрация
      <img width="974" height="427" alt="image" src="https://github.com/user-attachments/assets/ee72eaba-e529-44d0-adea-d8ba69a7a2a9" />

   - /auth/login
     Вход
     <img width="974" height="468" alt="image" src="https://github.com/user-attachments/assets/fbef121d-21c3-4be9-ad04-1fec871af175" />
     Неверный вход
     <img width="974" height="409" alt="image" src="https://github.com/user-attachments/assets/6e803a23-570a-4b2e-ace7-fd436e6ea229" />



