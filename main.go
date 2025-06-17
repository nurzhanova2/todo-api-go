package main

import (
	"github.com/nurzhanova/todo-app/routes"

	"github.com/gin-gonic/gin"
	"github.com/nurzhanova/todo-app/db"
)

func main() {
	db.InitDB() // подключаемся к БД

    r := gin.Default()

    routes.SetupRoutes(r) // подключаем маршруты

    r.Run(":8080") // запускаем сервер на порту 8080
}