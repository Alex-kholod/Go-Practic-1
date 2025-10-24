# Практическое занятие №7

## Тема: Подключение и работа с Redis (set/get, TTL). Реализация простого кэша

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz7 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
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
1. Запустили Redis локально в docker контейнере
```bash
docker run --name redis -p 6379:6379 redis
```
   
2. Запустили код и выполнили запросы
   - /set?key=test&value=hello
     <img width="974" height="64" alt="image" src="https://github.com/user-attachments/assets/855538c0-d87e-4695-9c4b-5481ec448e78" />

   - /get?key=test
     <img width="974" height="64" alt="image" src="https://github.com/user-attachments/assets/c0d90082-f7e3-46e5-849f-ba4ccd38f946" />

   - /ttl?key=test
     <img width="974" height="106" alt="image" src="https://github.com/user-attachments/assets/36f870ca-6c2c-48be-b3fb-abe65b53a537" />


## Описание
Redis (Remote Dictionary Server) — это высокопроизводительная in-memory система хранения данных, которая работает по принципу «ключ–значение».
Данные хранятся в оперативной памяти, что делает доступ к ним очень быстрым.

Встроенный механизм TTL (time-to-live) — время жизни ключей. Когда TTL истекает, ключ автоматически удаляется.
В текущем проекте TTL=10 секундам. Это означает, что через 10 секунд после запроса SET по запросу GET (с тем же ключом "test") мы уже не получим данных, так как они удалятся из памяти.
