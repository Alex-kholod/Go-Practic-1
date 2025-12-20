# Практическое занятие №15

## Тема: Unit-тестирование функций (testing, testify)

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz15 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
```
2. Переход в директорию проекта
```bash
cd Go-Practic-1
```

### Ход работы
1. Запуск всех тестов
```bash
go test ./...
```
<img width="650" height="198" alt="image" src="https://github.com/user-attachments/assets/34a8e32a-29b1-4f53-82a6-99dcbb8c6efc" />

2. Процент покрытия кода тестами
```bash
go test -cover ./...
```
<img width="974" height="187" alt="image" src="https://github.com/user-attachments/assets/52171ee7-b83c-4937-84ea-5de5f553fcf6" />

3. Бэнчмарк тест пакета mathx
```bash
go test -bench . ./internal/mathx
```
<img width="974" height="429" alt="image" src="https://github.com/user-attachments/assets/1971fc0e-4cd0-4380-b40f-b87c93a33c15" />

4. Бэнчмарк тест пакета stringsx
```bash
go test -bench . ./internal/stringsx
```
<img width="974" height="389" alt="image" src="https://github.com/user-attachments/assets/c35ad74f-64c4-4bf1-9732-3f7cf0e2adf5" />

5. Бэнчмарк тест пакета service
```bash
go test -bench . ./internal/service
```
<img width="974" height="300" alt="image" src="https://github.com/user-attachments/assets/fe8b6711-cb5e-4147-8461-1469270d4643" />

6. Вывод html отчета покрытия кода
   <img width="974" height="775" alt="image" src="https://github.com/user-attachments/assets/9cd147bb-9f4a-4550-b97d-9127f655c3de" />


### Структура проекта
<img width="539" height="742" alt="image" src="https://github.com/user-attachments/assets/a9bbf534-b5be-4947-89aa-89b8f6a49ec7" />
