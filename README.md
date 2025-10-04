# Практическое занятие №4

## Тема: Маршрутизация с chi. Создание небольшого CRUD-сервиса «Список задач».

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz4 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```
3. Запуск приложения
```bash
go run ./cmd/myapp
```

### Принцип работы
После запуска приложения можно использовать его endpoint's для взаимодействия с задачами. (CRUD)
#### Поддерживаемые endpoint's:
## 1. GET /
   #### Формат ответа:
   <img width="974" height="618" alt="image" src="https://github.com/user-attachments/assets/69f33b90-8039-423d-beb0-03a25298fc55" />

   
## 2. POST /
   #### Формат ответа:
   <img width="974" height="500" alt="image" src="https://github.com/user-attachments/assets/42cbb199-8ce9-4bd0-898f-0d4e2a99dae4" />


## 4. GET /{id}
   #### Формат ответа:
   <img width="974" height="544" alt="image" src="https://github.com/user-attachments/assets/95b63a5c-812b-40ec-b6d3-d5fff7b33af5" />

   
## 4. PUT /{id}
   #### Формат ответа:
   <img width="974" height="519" alt="image" src="https://github.com/user-attachments/assets/1be2cbfd-cd29-4ab4-897e-caaae11a0d97" />


## 4. DELETE /{id}
   #### Формат ответа:
   <img width="974" height="418" alt="image" src="https://github.com/user-attachments/assets/1af5c972-58a6-4bc5-82f7-1e61a9d431e3" />



### Дополнительно
1. Валидация длины названия задачи.
   <img width="974" height="435" alt="image" src="https://github.com/user-attachments/assets/400b5d71-6812-4bce-a5b1-e1d4a8219dda" />

2. Пагинация
   <img width="974" height="618" alt="image" src="https://github.com/user-attachments/assets/03bcb3f0-bfdf-4363-99ad-a485a0cf2987" />
   <img width="974" height="614" alt="image" src="https://github.com/user-attachments/assets/791f3bb0-f6a4-4f08-b0b9-eab40a91ac77" />

3. Фильтр по параметру "done=true"
   <img width="974" height="540" alt="image" src="https://github.com/user-attachments/assets/9bdfcaeb-3b8f-4250-9221-533e96173826" />

4. Чтение и запись в файл
   <img width="974" height="772" alt="image" src="https://github.com/user-attachments/assets/b30490b4-4d7c-43d6-af77-ece48a966c52" />

5. Версионность апи
   <img width="974" height="585" alt="image" src="https://github.com/user-attachments/assets/edbd7cb3-07f5-4e13-bb6b-5594e5ca995b" />


