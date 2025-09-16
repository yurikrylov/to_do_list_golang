// baseURL/part_6/6.1/golang/todo/main.go
package main

import (
	"time"

	"github.com/yurikrylov/to_do_list_golang/db"
	"github.com/yurikrylov/to_do_list_golang/menu"
)

func main() {
	// Создаем репозиторий
	rep := db.NewSQLiteRepository()
	// Создаем отложенное закрытие соединения
	defer rep.Close()
	// Бесконечный цикл
	for {
		menu.CreateMenu(rep)
		time.Sleep(2 * time.Second)
	}
}
