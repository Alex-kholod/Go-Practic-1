# Практическое занятие №10

## Тема: JWT-аутентификация: создание и проверка токенов. Middleware для авторизации

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz10 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./cmd/server
```

### Ход работы
1. Реализовали аунтефикацию и авторизацию с помощью JWT-токенов с использованием алгоритма HS256 (HMAC-SHA256).
   
   Сделали refresh механику обновления токенов по окначанию их срока действия.

   Access-токены с TTL 15 минут и Refresh-токены с TTL 7 дней.

   In-memory blacklist для отозванных токенов.

2. Запустили код и выполнили запросы на регистрацию и авторизацию
3. Логин — получить токен
   - /api/v1/login
     <img width="974" height="431" alt="image" src="https://github.com/user-attachments/assets/212a0765-b8b7-42e8-a750-96b0ac9cd92f" />

4. Доступ к защищённым ручкам
   - /api/v1/me
     <img width="974" height="601" alt="image" src="https://github.com/user-attachments/assets/8389c218-9e7a-474f-9a2c-233622455b8a" />

   - /api/v1/admin/stats
     <img width="974" height="566" alt="image" src="https://github.com/user-attachments/assets/3bc8068e-61d8-45a4-9803-90e2075fa054" />

5. Авторизация с ролью user
   - /api/v1/login
     <img width="974" height="658" alt="image" src="https://github.com/user-attachments/assets/24a6af79-7bbe-4ba4-b8d4-b3c1fbc6ad57" />

   Пробуем получить статистику админов (с ролью user)
   - /api/v1/admin/stats
     <img width="974" height="560" alt="image" src="https://github.com/user-attachments/assets/bc602d24-6d7b-470f-a0e6-97e9909ec495" />

6. Новая механика с refresh и access токенами
   - Логин под admin - получение токенов
     <img width="974" height="607" alt="image" src="https://github.com/user-attachments/assets/a614aa44-9a29-432e-9a09-1918a546ed57" />
     
   - Получение статистики для админов
     <img width="974" height="548" alt="image" src="https://github.com/user-attachments/assets/4ab5a8de-da7f-4776-846a-02a0a4c27df8" />
     
   - Логин под user - получение токенов
     <img width="974" height="614" alt="image" src="https://github.com/user-attachments/assets/0388a266-fc6d-4a98-b518-cbcbdcb51ce9" />

   - Получение информации о себе
     <img width="974" height="562" alt="image" src="https://github.com/user-attachments/assets/1169d752-9148-4d29-a4bd-9cbf99ebd56b" />

   - Попытка получить статиски для админов
     <img width="974" height="505" alt="image" src="https://github.com/user-attachments/assets/4ac97ec9-ac09-433f-ae19-8fc4b97a175e" />

7. Обновление токенов
   - /api/v1/refresh
     <img width="974" height="550" alt="image" src="https://github.com/user-attachments/assets/3fcf8e9c-bb2a-41b2-85ba-e8c4ac2ab58d" />

8. GET /api/v1/me запрос с новыми токеном под user
   <img width="974" height="550" alt="image" src="https://github.com/user-attachments/assets/090ff3d7-91b9-4126-85cf-fb004119bb16" />

   

