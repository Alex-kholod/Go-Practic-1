Практическое занятие №2 
Тема: Структура Go-проекта
Студент: Холодков А.Д. Группа ЭФМО-01-25

Запуск проекта
1. Склонировать ветку репозитория
   git clone -b pz2 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
2. Запустить через терминал командой: 
   "go run ./cmd/myapp"

Принцип работы
После запуска приложения можно использовать его endpoint для взаимодействия.
Поддерживаемые endpoint's:
1. /
   Формат ответа:
   ![alt text](image.png)
2. /ping
   Формат ответа:
   a) без передачи заголовка
   ![alt text](image-1.png)
   b) с передачей заголовка "X-Request-Id"
   ![alt text](image-2.png)

3. /fail
   Формат ответа:
   ![alt text](image-3.png)

Структура проекта:
![alt text](image-4.png)