package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandlerHello(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello string")
}

func HandlerPing(ctx *gin.Context) {
    pong, err := RedisClient.Ping()

    if err != nil {
        ctx.String(http.StatusInternalServerError, err.Error())
        return
    }

    ctx.String(http.StatusOK, pong)
}

