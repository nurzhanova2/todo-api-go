package handlers

import (
    "net/http"
    "github.com/nurzhanova/todo-app/db"
    "github.com/nurzhanova/todo-app/models"

    "github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
    var tasks []models.Task

    err := db.DB.Select(&tasks, "SELECT * FROM tasks")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить задачи"})
        return
    }

    c.JSON(http.StatusOK, tasks)
}

func CreateTask(c *gin.Context) {
    var task models.Task

    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
        return
    }

    query := `INSERT INTO tasks (title, description, completed) 
              VALUES ($1, $2, $3) RETURNING id`

    err := db.DB.QueryRow(query, task.Title, task.Description, task.Completed).Scan(&task.ID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании задачи"})
        return
    }

    c.JSON(http.StatusCreated, task)
}

func UpdateTask(c *gin.Context) {
    id := c.Param("id") 
    var task models.Task

    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный JSON"})
        return
    }

    query := `
        UPDATE tasks 
        SET title = $1, description = $2, completed = $3 
        WHERE id = $4
    `

    _, err := db.DB.Exec(query, task.Title, task.Description, task.Completed, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить задачу"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Задача обновлена"})
}

func DeleteTask(c *gin.Context) {
    id := c.Param("id")

    query := `DELETE FROM tasks WHERE id = $1`
    _, err := db.DB.Exec(query, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить задачу"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}

func GetTaskByID(c *gin.Context) {
    id := c.Param("id")

    var task models.Task

    query := `SELECT * FROM tasks WHERE id = $1`
    err := db.DB.Get(&task, query, id)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
        return
    }

    c.JSON(http.StatusOK, task)
}
