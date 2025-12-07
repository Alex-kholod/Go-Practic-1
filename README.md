# Практическое занятие №13

## Тема: Профилирование Go-приложения (pprof). Измерение времени работы функций

**Студент:** Холодков А.Д.  
**Группа:** ЭФМО-01-25

### Установка и запуск
1. Клонирование репозитория (только ветка текущей практики)
```bash
git clone -b pz13 --single-branch https://github.com/Alex-kholod/Go-Practic-1.git
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
1. Подключили профилировщик pprof. Открыли его Web-интерфейс
   <img width="974" height="460" alt="image" src="https://github.com/user-attachments/assets/eb996333-bbe1-489e-a47e-8c850f662198" />

3. Запустили анализ использоваения CPU на протяжении 30 секунд.
   - Top по потреблению CPU
     <img width="974" height="289" alt="image" src="https://github.com/user-attachments/assets/d0f15cb4-1bd2-4da9-8189-d017d15f56c8" />

   - «Горящие» строки
     <img width="974" height="179" alt="image" src="https://github.com/user-attachments/assets/0955c4ba-bc5d-4a9e-9e1c-e3491e68350d" />

   - Граф вызовов
     <img width="974" height="670" alt="image" src="https://github.com/user-attachments/assets/f253ccb2-c676-4f6c-8fb6-99a4785e0bd1" />

4. Визуализация использования памяти
   <img width="974" height="605" alt="image" src="https://github.com/user-attachments/assets/66325d61-b4b4-41bd-b693-a906add202be" />

   - alloc_space - вся выделенная память за всё время
     <img width="974" height="495" alt="image" src="https://github.com/user-attachments/assets/6c78f54f-197e-4d52-aae7-be7f769655b2" />

   - inuse_space - используемая память в данный момент
     <img width="974" height="496" alt="image" src="https://github.com/user-attachments/assets/0ea04278-4fee-42f0-aa86-439a8539a53d" />

6. Добавили свой таймер для функции обработчика с выводом длительности каждого запроса.
   <img width="908" height="302" alt="image" src="https://github.com/user-attachments/assets/845ff3d3-ba06-4e01-bf3c-52141b3c3507" />

8. Создали бенчмарк тест и запустили до оптимизации.
   <img width="974" height="244" alt="image" src="https://github.com/user-attachments/assets/95b8ec43-edb9-4e81-ac13-21eba508d06c" />

   - 274 – количество итераций
   - ns/op - Время выполнения на одну операцию (наносекунды)
   - B/op - Память на одну операцию (байты)
   - allocs/op - Количество аллокаций на операцию (сколько раз выделялась память)
10. Заменили рекурсию на циклический перебор и сделали новый бенчмарк на тестирование скорости.
    <img width="974" height="229" alt="image" src="https://github.com/user-attachments/assets/8707a5a4-9a2d-4433-8c94-17e7bdd5b798" />

   
### Итог: Количество итераций увеличилось, но при этом время выполнения на одну операцию уменьшилось, следовательно выросла скорость работы программы.



