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
2. Процент покрытия кода тестами
```bash
go test -cover ./...
```
3. Бэнчмарк тест пакета mathx
```bash
go test -bench . ./internal/mathx
```
4. Бэнчмарк тест пакета stringsx
```bash
go test -bench . ./internal/stringsx
```
5. Бэнчмарк тест пакета service
```bash
go test -bench . ./internal/service
```
6. Вывод html отчета покрытия кода

### Структура проекта