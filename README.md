# Практическое занятие №12

## Тема: Подключение Swagger/OpenAPI. Автоматическая генерация документации

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz12 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
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
1. Подключили автогенерацию документации к проекту API для заметок.
2. Для автогенерации добавили аннотации к функциям в коде.
3. Добавили endpoint GET /docs для просмотра документации.
   <img width="974" height="493" alt="image" src="https://github.com/user-attachments/assets/99a8bc32-4bea-458a-ad96-9bc384a0d073" />

5. Вызываем метод создания заметки через документацию:
   <img width="974" height="493" alt="image" src="https://github.com/user-attachments/assets/09055e88-647c-4e86-8d3e-5efad8d69969" />

6. Метод получения списка заметок через документацию:
   <img width="974" height="426" alt="image" src="https://github.com/user-attachments/assets/2b621059-699f-4aa8-a217-815606f4e271" />

7. DTO (Модели данных)
   <img width="974" height="407" alt="image" src="https://github.com/user-attachments/assets/b429ce99-060f-43c2-87c4-50101c313e30" />


   
### Структура проекта
<img width="538" height="1028" alt="image" src="https://github.com/user-attachments/assets/75fd8204-d05f-4b49-9005-84355856bb1b" />

