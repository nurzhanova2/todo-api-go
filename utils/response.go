package utils

import (

    "github.com/gin-gonic/gin"
)

type ApiResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, statusCode int, data interface{}) {
    c.JSON(statusCode, ApiResponse{
        Success: true,
        Data:    data,
    })
}

func Error(c *gin.Context, statusCode int, message string) {
    c.JSON(statusCode, ApiResponse{
        Success: false,
        Error:   message,
    })
}
