package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	db, err := openDB(dsn)
	if err != nil {
		log.Fatalf("openDB error: %v", err)
	}
	defer db.Close()

	repo := NewRepo(db)

	// 1) Вставим пару задач
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	titles := []string{"Сделать ПЗ №5", "Купить кофе", "Проверить отчёты"}
	for _, title := range titles {
		id, err := repo.CreateTask(ctx, title)
		if err != nil {
			log.Fatalf("CreateTask error: %v", err)
		}
		log.Printf("Inserted task id=%d (%s)", id, title)
	}

	// 2) Прочитаем список задач
	ctxList, cancelList := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelList()

	tasks, err := repo.ListTasks(ctxList)
	if err != nil {
		log.Fatalf("ListTasks error: %v", err)
	}

	// 3) Напечатаем
	fmt.Println("=== Tasks ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	// 4) Только выполненные задачи
	ctxListDone, cancelListDone := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelListDone()

	tasks, err = repo.ListDone(ctxListDone, true)
	if err != nil {
		log.Fatalf("ListDone error: %v", err)
	}

	fmt.Println("=== Tasks DONE ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	// 5) Только не выполненные задачи
	ctxListNoDone, cancelListNoDone := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelListNoDone()

	tasks, err = repo.ListDone(ctxListNoDone, false)
	if err != nil {
		log.Fatalf("ListDone error: %v", err)
	}

	fmt.Println("=== Tasks NO DONE ===")
	for _, t := range tasks {
		fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
			t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))
	}

	// 6) Конкретная задача по ID
	ctxTask, cancelTask := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelTask()

	t, err := repo.FindByID(ctxTask, 1)
	if err != nil {
		log.Fatalf("FindByID error: %v", err)
	}

	fmt.Println("=== Tasks By ID ===")
	fmt.Printf("#%d | %-24s | done=%-5v | %s\n",
		t.ID, t.Title, t.Done, t.CreatedAt.Format(time.RFC3339))

	// 7) Массовая вставка через транзакцию
	ctxListMany, cancelListMany := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelListMany()

	titles = []string{"Задача 1", "Задача 2", "Задача 3"}
	tasksIds, err := repo.CreateMany(ctxListMany, titles)
	if err != nil {
		log.Fatalf("CreateMany error: %v", err)
	}

	fmt.Println("=== Tasks Insert Many ===")
	for _, id := range tasksIds {
		fmt.Printf("#%d", id)
	}

}
