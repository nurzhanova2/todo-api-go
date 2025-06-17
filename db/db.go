package db

import (
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Ошибка загрузки .env файла: %v", err)
    }

    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    sslmode := os.Getenv("SSL_MODE")

    dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
        user, password, host, port, dbname, sslmode)

    DB, err = sqlx.Connect("postgres", dbURL)
    if err != nil {
        log.Fatalf("Ошибка подключения к базе данных: %v", err)
    }

    fmt.Println("Успешное подключение к базе данных!")

    createTable()
}

func createTable() {
    schema := `
    CREATE TABLE IF NOT EXISTS tasks (
        id SERIAL PRIMARY KEY,
        title TEXT NOT NULL,
        description TEXT,
        completed BOOLEAN DEFAULT false
    );`

    DB.MustExec(schema)
    fmt.Println("Таблица tasks готова.")
}